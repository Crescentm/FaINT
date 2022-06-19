package parameterType

import (
	"Fuzzer/src/random"
	"math/big"
)

type BigUintType struct {
	ParameterType
	MaxValue big.Int
	MinValue big.Int
}

func NewBigUintType(param *ParameterType) *BigUintType {
	var size = param.Size
	var maxValue, minValue big.Int
	maxValue.Lsh(big.NewInt(1), uint(size))
	minValue.SetUint64(0)
	return &BigUintType{ParameterType: *param, MaxValue: maxValue, MinValue: minValue}
}

func (t *BigUintType) Generate() interface{} {
	return random.RandBigInt(&t.MinValue, &t.MaxValue)
}
