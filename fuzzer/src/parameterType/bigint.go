package parameterType

import (
	"Fuzzer/src/random"
	"math/big"
)

type BigIntType struct {
	ParameterType
	MaxValue big.Int
	MinValue big.Int
}

func NewBigIntType(param *ParameterType) *BigIntType {
	var size = param.Size
	var maxValue, minValue big.Int
	maxValue.Lsh(big.NewInt(1), uint(size-1))
	//minValue.Lsh(big.NewInt(-1), uint(size-2))
	minValue.SetInt64(0)
	return &BigIntType{ParameterType: *param, MaxValue: maxValue, MinValue: minValue}
}

func (t *BigIntType) Generate() interface{} {
	return random.RandBigInt(&t.MinValue, &t.MaxValue)
}
