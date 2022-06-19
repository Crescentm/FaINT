// Copyright 2020 ConsenSys Software Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by consensys/gnark-crypto DO NOT EDIT

package fft

import (
	"math/bits"
	"runtime"

	"github.com/consensys/gnark-crypto/internal/parallel"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
)

// Decimation is used in the FFT call to select decimation in time or in frequency
type Decimation uint8

const (
	DIT Decimation = iota
	DIF
)

// parallelize threshold for a single butterfly op, if the fft stage is not parallelized already
const butterflyThreshold = 16

// FFT computes (recursively) the discrete Fourier transform of a and stores the result in a
// if decimation == DIT (decimation in time), the input must be in bit-reversed order
// if decimation == DIF (decimation in frequency), the output will be in bit-reversed order
// coset sets the shift of the fft (0 = no shift, standard fft)
// len(a) must be a power of 2, and w must be a len(a)th root of unity in field F.
//
// example:
// -------
// domain := NewDomain(m, 2) -->  contains precomputed data for Z/mZ, and Z/4mZ
// FFT(pol, DIT, 1) --> evaluates pol on the coset 1 in (Z/4mZ)/(Z/mZ)
func (domain *Domain) FFT(a []fr.Element, decimation Decimation, coset uint64) {

	numCPU := uint64(runtime.NumCPU())

	if coset != 0 {
		if decimation == DIT {
			BitReverse(domain.CosetTable[coset-1])
		}
		parallel.Execute(len(a), func(start, end int) {
			for i := start; i < end; i++ {
				a[i].Mul(&a[i], &domain.CosetTable[coset-1][i])
			}
		})
		// put it back as we found it
		if decimation == DIT {
			BitReverse(domain.CosetTable[coset-1])
		}
	}

	// find the stage where we should stop spawning go routines in our recursive calls
	// (ie when we have as many go routines running as we have available CPUs)
	maxSplits := bits.TrailingZeros64(nextPowerOfTwo(numCPU))
	if numCPU <= 1 {
		maxSplits = -1
	}

	switch decimation {
	case DIF:
		difFFT(a, domain.Twiddles, 0, maxSplits, nil)
	case DIT:
		ditFFT(a, domain.Twiddles, 0, maxSplits, nil)
	default:
		panic("not implemented")
	}
}

// FFTInverse computes (recursively) the inverse discrete Fourier transform of a and stores the result in a
// if decimation == DIT (decimation in time), the input must be in bit-reversed order
// if decimation == DIF (decimation in frequency), the output will be in bit-reversed order
// coset sets the shift of the fft (0 = no shift, standard fft)
// len(a) must be a power of 2, and w must be a len(a)th root of unity in field F.
func (domain *Domain) FFTInverse(a []fr.Element, decimation Decimation, coset uint64) {

	numCPU := uint64(runtime.NumCPU())

	// find the stage where we should stop spawning go routines in our recursive calls
	// (ie when we have as many go routines running as we have available CPUs)
	maxSplits := bits.TrailingZeros64(nextPowerOfTwo(numCPU))
	if numCPU <= 1 {
		maxSplits = -1
	}
	switch decimation {
	case DIF:
		difFFT(a, domain.TwiddlesInv, 0, maxSplits, nil)
	case DIT:
		ditFFT(a, domain.TwiddlesInv, 0, maxSplits, nil)
	default:
		panic("not implemented")
	}

	// scale by CardinalityInv (+ cosetTableInv is coset!=0)
	if coset != 0 {
		if decimation == DIF {
			BitReverse(domain.CosetTableInv[coset-1])
		}
		parallel.Execute(len(a), func(start, end int) {
			for i := start; i < end; i++ {
				a[i].Mul(&a[i], &domain.CosetTableInv[coset-1][i]).
					Mul(&a[i], &domain.CardinalityInv)
			}
		})
		// put it back as we found it
		if decimation == DIF {
			BitReverse(domain.CosetTableInv[coset-1])
		}
	} else {
		parallel.Execute(len(a), func(start, end int) {
			for i := start; i < end; i++ {
				a[i].Mul(&a[i], &domain.CardinalityInv)
			}
		})
	}
}

func difFFT(a []fr.Element, twiddles [][]fr.Element, stage, maxSplits int, chDone chan struct{}) {
	if chDone != nil {
		defer func() {
			chDone <- struct{}{}
		}()
	}
	n := len(a)
	if n == 1 {
		return
	}
	m := n >> 1

	// if stage < maxSplits, we parallelize this butterfly
	// but we have only numCPU / stage cpus available
	if (m > butterflyThreshold) && (stage < maxSplits) {
		// 1 << stage == estimated used CPUs
		numCPU := runtime.NumCPU() / (1 << (stage))
		parallel.Execute(m, func(start, end int) {
			var t fr.Element
			for i := start; i < end; i++ {
				t = a[i]
				a[i].Add(&a[i], &a[i+m])

				a[i+m].
					Sub(&t, &a[i+m]).
					Mul(&a[i+m], &twiddles[stage][i])
			}
		}, numCPU)
	} else {
		var t fr.Element

		// i == 0
		t = a[0]
		a[0].Add(&a[0], &a[m])
		a[m].Sub(&t, &a[m])

		for i := 1; i < m; i++ {
			t = a[i]
			a[i].Add(&a[i], &a[i+m])

			a[i+m].
				Sub(&t, &a[i+m]).
				Mul(&a[i+m], &twiddles[stage][i])
		}
	}

	if m == 1 {
		return
	}

	nextStage := stage + 1
	if stage < maxSplits {
		chDone := make(chan struct{}, 1)
		go difFFT(a[m:n], twiddles, nextStage, maxSplits, chDone)
		difFFT(a[0:m], twiddles, nextStage, maxSplits, nil)
		<-chDone
	} else {
		difFFT(a[0:m], twiddles, nextStage, maxSplits, nil)
		difFFT(a[m:n], twiddles, nextStage, maxSplits, nil)
	}
}

func ditFFT(a []fr.Element, twiddles [][]fr.Element, stage, maxSplits int, chDone chan struct{}) {
	if chDone != nil {
		defer func() {
			chDone <- struct{}{}
		}()
	}
	n := len(a)
	if n == 1 {
		return
	}
	m := n >> 1

	nextStage := stage + 1

	if stage < maxSplits {
		// that's the only time we fire go routines
		chDone := make(chan struct{}, 1)
		go ditFFT(a[m:], twiddles, nextStage, maxSplits, chDone)
		ditFFT(a[0:m], twiddles, nextStage, maxSplits, nil)
		<-chDone
	} else {
		ditFFT(a[0:m], twiddles, nextStage, maxSplits, nil)
		ditFFT(a[m:n], twiddles, nextStage, maxSplits, nil)

	}

	// if stage < maxSplits, we parallelize this butterfly
	// but we have only numCPU / stage cpus available
	if (m > butterflyThreshold) && (stage < maxSplits) {
		// 1 << stage == estimated used CPUs
		numCPU := runtime.NumCPU() / (1 << (stage))
		parallel.Execute(m, func(start, end int) {
			var t, tm fr.Element
			for k := start; k < end; k++ {
				t = a[k]
				tm.Mul(&a[k+m], &twiddles[stage][k])
				a[k].Add(&a[k], &tm)
				a[k+m].Sub(&t, &tm)
			}
		}, numCPU)

	} else {
		var t, tm fr.Element
		// k == 0
		// wPow == 1
		t = a[0]
		a[0].Add(&a[0], &a[m])
		a[m].Sub(&t, &a[m])

		for k := 1; k < m; k++ {
			t = a[k]
			tm.Mul(&a[k+m], &twiddles[stage][k])
			a[k].Add(&a[k], &tm)
			a[k+m].Sub(&t, &tm)
		}
	}
}

// BitReverse applies the bit-reversal permutation to a.
// len(a) must be a power of 2 (as in every single function in this file)
func BitReverse(a []fr.Element) {
	n := uint64(len(a))
	nn := uint64(64 - bits.TrailingZeros64(n))

	for i := uint64(0); i < n; i++ {
		irev := bits.Reverse64(i) >> nn
		if irev > i {
			a[i], a[irev] = a[irev], a[i]
		}
	}
}
