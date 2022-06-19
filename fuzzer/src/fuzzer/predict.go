package fuzzer

import (
	"Fuzzer/src/parameterType"
	. "Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"math/big"
)

func Predict(input INPUTDATA, cost COST, inputNew INPUTDATA, costNew COST, choice int, inputType INPUTTYPE) INPUTDATA {
	var predictedInput = make(INPUTDATA, len(input))
	copy(predictedInput, input)

	parameter := parameterType.NewParameterType(&inputType[choice])
	cost1, cost2 := tools.PickDiffCost(cost, costNew)
	input1 := big.NewInt(0)
	input2 := big.NewInt(0)

	if parameter.TypeName[:3] != "int" && parameter.TypeName[:4] != "uint" {
		panic("wrong predict type")
	}

	if parameter.Size == 8 || parameter.Size == 16 || parameter.Size == 32 || parameter.Size == 64 {
		if parameter.TypeName[:3] == "int" {
			input1 = big.NewInt(input[choice].(int64))
			input2 = big.NewInt(inputNew[choice].(int64))
		} else if parameter.TypeName[:4] == "uint" {
			input1.SetUint64(input[choice].(uint64))
			input2.SetUint64(inputNew[choice].(uint64))
		}
	} else {
		input1 = input[choice].(*big.Int)
		input2 = inputNew[choice].(*big.Int)
	}

	//println("==============")
	//fmt.Println(input1)
	//fmt.Println(input2)
	//println("--------------")
	if cost1 == nil || cost2 == nil {
		//fmt.Println("not diff cost")
		//fmt.Println("-------------------")
		//fmt.Println(cost)
		//fmt.Println(costNew)
		//fmt.Println("-------------------")
		return predictedInput
	}
	argNew := big.NewInt(0)
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	mul1 := big.NewInt(0)
	mul2 := big.NewInt(0)
	mul1.Mul(cost1, input2)
	mul2.Mul(cost2, input1)

	numerator.Sub(mul1, mul2)
	//fmt.Println(mul1)
	//fmt.Println(mul2)
	//fmt.Println(numerator)
	//println("--------------")
	denominator.Sub(cost1, cost2)
	//fmt.Println(cost1)
	//fmt.Println(cost2)
	//fmt.Println(denominator)
	//println("--------------")

	argNew.Div(numerator, denominator)
	//fmt.Println("new arg:", argNew)
	//println("==============")

	if parameter.Size == 8 || parameter.Size == 16 || parameter.Size == 32 || parameter.Size == 64 {
		if parameter.TypeName[:3] == "int" {
			switch parameter.Size {
			case 8:
				predictedInput[choice] = int8(argNew.Int64())
			case 16:
				predictedInput[choice] = int16(argNew.Int64())
			case 32:
				predictedInput[choice] = int32(argNew.Int64())
			case 64:
				predictedInput[choice] = argNew.Int64()
			}
		} else if parameter.TypeName[:4] == "uint" {
			switch parameter.Size {
			case 8:
				predictedInput[choice] = uint8(argNew.Int64())
			case 16:
				predictedInput[choice] = uint16(argNew.Int64())
			case 32:
				predictedInput[choice] = uint32(argNew.Int64())
			case 64:
				predictedInput[choice] = uint64(argNew.Int64())
			}
		}
	} else {
		predictedInput[choice] = argNew
	}

	return predictedInput
}
