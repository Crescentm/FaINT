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

package amd64

func (f *FFAmd64) generateDouble() {
	// func header
	f.Comment("double(res, x *Element)")
	stackSize := f.StackSize(f.NbWords*2, 0, 0)
	registers := f.FnHeader("double", stackSize, 16)
	defer f.AssertCleanStack(stackSize, 0)

	// registers
	x := registers.Pop()
	t := registers.PopN(f.NbWords)

	f.MOVQ("x+8(FP)", x)
	f.Mov(x, t)
	registers.Push(x)

	f.Add(t, t)
	f.Reduce(&registers, t)

	r := registers.Pop()
	f.MOVQ("res+0(FP)", r)
	f.Mov(t, r)

	f.RET()
}
