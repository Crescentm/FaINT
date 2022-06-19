package tools

import (
	"Fuzzer/src/parameterType"
	"Fuzzer/src/pidType"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func RandomInput(method abi.Method) pidType.INPUTDATA {
	var inputList pidType.INPUTDATA
	for _, arg := range method.Inputs {
		argumentType := arg.Type
		parameter := parameterType.NewParameterType(&argumentType)
		switch {
		case parameter.TypeName[:4] == "int8", parameter.TypeName[:5] == "int16", parameter.TypeName[:5] == "int32", parameter.TypeName[:5] == "int64":
			varInt := parameterType.NewBuildinIntType(parameter)
			inputList = append(inputList, varInt.Generate())
			break
		case parameter.TypeName[:5] == "uint8", parameter.TypeName[:6] == "uint16", parameter.TypeName[:6] == "uint32", parameter.TypeName[:6] == "uint64":
			varUint := parameterType.NewBuildinUintType(parameter)
			inputList = append(inputList, varUint.Generate())
			break
		case parameter.TypeName[:3] == "int":
			varInt := parameterType.NewBigIntType(parameter)
			inputList = append(inputList, varInt.Generate())
			break
		case parameter.TypeName[:4] == "uint":
			varUint := parameterType.NewBigUintType(parameter)
			inputList = append(inputList, varUint.Generate())
			break
		case parameter.TypeName[:4] == "bool":
			varBool := parameterType.NewBoolType(parameter)
			inputList = append(inputList, varBool.Generate())
			break
		case parameter.TypeName[:4] == "byte":
			varByte := parameterType.NewByteType(parameter)
			inputList = append(inputList, varByte.Generate())
			break
		case parameter.TypeName[:6] == "string":
			varString := parameterType.NewStringType(parameter, 24)
			inputList = append(inputList, varString.Generate())
			break
		default:
			continue
		}
	}
	return inputList
}
