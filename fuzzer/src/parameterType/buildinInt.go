package parameterType

import (
	"Fuzzer/src/random"
	"encoding/binary"
	"github.com/adam-lavrik/go-imath/i16"
	"github.com/adam-lavrik/go-imath/i32"
	"github.com/adam-lavrik/go-imath/i64"
	"github.com/adam-lavrik/go-imath/i8"
)

type BuildinIntType struct {
	ParameterType
	Size int
}

func NewBuildinIntType(param *ParameterType) *BuildinIntType {
	var size = param.Size
	return &BuildinIntType{ParameterType: *param, Size: size}
}

func (t *BuildinIntType) Generate() interface{} {
	buf := random.RandByte(t.Size)
	switch t.Size {
	case 8:
		return i8.Abs(int8(binary.LittleEndian.Uint64(buf)))
	case 16:
		return i16.Abs(int16(binary.LittleEndian.Uint64(buf)))
	case 32:
		return i32.Abs(int32(binary.LittleEndian.Uint64(buf)))
	case 64:
		return i64.Abs(int64(binary.LittleEndian.Uint64(buf)))
	default:
		return nil
	}
}
