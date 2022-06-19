package fuzzer

import (
	"Fuzzer/src/parameterType"
	. "Fuzzer/src/pidType"
	"errors"
	"math/rand"
	"time"
)

func FuzzInput(input INPUTDATA, inputType INPUTTYPE) (INPUTDATA, int, error) {
	if input == nil {
		return input, 0, errors.New("no input")
	}
	inputNew := make(INPUTDATA, len(input))
	copy(inputNew, input)

	//which one input to variate
	var parameter *parameterType.ParameterType
	var choice int

	for {
		choice = RandomChoose(len(input))
		parameter = parameterType.NewParameterType(&inputType[choice])

		if parameter.TypeName[:3] == "int" || parameter.TypeName[:4] == "uint" {
			break
		}
	}

	switch {
	case parameter.TypeName[:4] == "int8", parameter.TypeName[:5] == "int16", parameter.TypeName[:5] == "int32", parameter.TypeName[:5] == "int64":
		varBuildinInt := parameterType.NewBuildinIntType(parameter)
		inputNew[choice] = varBuildinInt.Generate()
		break
	case parameter.TypeName[:5] == "uint8", parameter.TypeName[:6] == "uint16", parameter.TypeName[:6] == "uint32", parameter.TypeName[:6] == "uint64":
		varBuildinUint := parameterType.NewBuildinUintType(parameter)
		inputNew[choice] = varBuildinUint.Generate()
		break
	case parameter.TypeName[:3] == "int":
		varInt := parameterType.NewBigIntType(parameter)
		inputNew[choice] = varInt.Generate()
		break
	case parameter.TypeName[:4] == "uint":
		varUint := parameterType.NewBigUintType(parameter)
		inputNew[choice] = varUint.Generate()
		break
	}

	return inputNew, choice, nil
}

func RandomChoose(maxNum int) int {
	rand.Seed(time.Now().Unix())
	v := rand.Intn(maxNum)
	return v
}
