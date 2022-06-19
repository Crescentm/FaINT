package fuzzer

import (
	. "Fuzzer/src/contract"
	. "Fuzzer/src/pidType"
	. "Fuzzer/src/tools"
	"fmt"
	"math/big"
)

var runTime = 0

func FuctionFuzz(fuctionName string, seeds any) {
	var err error
	methodObject := AbiObject.Methods[fuctionName]

	inputType := InputParse(methodObject)

	pids := RunSeeds(seeds, methodObject, inputType)

	for !Interrupted() {
		inputData, cost := PickInput(pids)
		energy := 0
		maxEnergy := AssignEnergy(inputData)

		var predictedInputData INPUTDATA

		for energy < maxEnergy || predictedInputData != nil {
			var inputDataNew INPUTDATA
			var choice = 0

			if predictedInputData != nil {
				inputDataNew = make(INPUTDATA, len(predictedInputData))
				copy(inputDataNew, predictedInputData)
				predictedInputData = nil
			} else {
				inputDataNew, choice, err = FuzzInput(inputData, inputType)
				if err != nil {
					panic("can not fuzzer input")
				}
			}

			pidCost, pidPath, status := SendInput(inputDataNew, fuctionName, ContractAddress, AbiObject, CodeId)
			pidNew := Constructor(inputDataNew, inputType, pidCost, pidPath, status)

			pidPathid := GeneratePathID(pidPath)
			input1 := inputDataNew[0].(*big.Int)
			if IsNew(pidPathid, pids) {
				energy = 0
				AddPid(pidNew, pidPathid, pids)
				fmt.Println("========new path========")
				fmt.Println("path:  ", pidPath)
				fmt.Println("input: ", input1)
				fmt.Println("status:", status)
				fmt.Println("========================")
			}

			if energy < maxEnergy {
				costNew := pidNew.Cost
				predictedInputData = Predict(inputData, cost, inputDataNew, costNew, choice, inputType)
			}

			energy = energy + 1
		}
		runTime = runTime + 1
	}
	fmt.Println("===========path sum=============")
	println("path num", len(pids))
	fmt.Println()
	for _, pid := range pids {
		fmt.Println("path :", pid.Path)
	}
	fmt.Println("===========path sum=============")

	fmt.Println("===========input-cost=============")
	for _, pid := range pids {
		fmt.Println("-----------------------")
		fmt.Println("input :", pid.InputData[0].(*big.Int))
		fmt.Println("cost  :")
		for id, cost := range pid.Cost {
			fmt.Println(id, cost)
		}
	}
	fmt.Println("===========input-cost=============")
}
