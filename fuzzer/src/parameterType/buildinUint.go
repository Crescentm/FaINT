package parameterType

import (
	"Fuzzer/src/random"
	"encoding/binary"
)

type BuildinUintType struct {
	ParameterType
	Size int
}

func NewBuildinUintType(param *ParameterType) *BuildinUintType {
	var size = param.Size
	return &BuildinUintType{ParameterType: *param, Size: size}
}

func (t *BuildinUintType) Generate() interface{} {
	buf := random.RandByte(t.Size)
	switch t.Size {
	case 8:
		return uint8(binary.LittleEndian.Uint64(buf))
	case 16:
		return uint16(binary.LittleEndian.Uint64(buf))
	case 32:
		return uint32(binary.LittleEndian.Uint64(buf))
	case 64:
		return binary.LittleEndian.Uint64(buf)
	default:
		return nil
	}
}
