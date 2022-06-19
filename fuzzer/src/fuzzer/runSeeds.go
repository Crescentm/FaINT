package fuzzer

import (
	"Fuzzer/src/contract"
	"Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

func RunSeeds(seeds any, method abi.Method, inputType pidType.INPUTTYPE) pidType.PIDS {
	var pids = make(pidType.PIDS)

	for i := 0; i < 4; i++ {
		inputList := tools.RandomInput(method)
		cost, pathInt, status := contract.SendInput(inputList, method.Name, ContractAddress, AbiObject, CodeId)

		pathid := pidType.GeneratePathID(pathInt)
		pidNew := pidType.Constructor(inputList, inputType, cost, pathInt, status)

		pids[pathid] = pidNew
	}
	for _, pid := range pids {
		input := pid.InputData[0].(*big.Int)

		fmt.Println("========new path========")
		fmt.Println("path:  ", pid.Path)
		fmt.Println("input: ", input)
		fmt.Println("status:", pid.Status)
		fmt.Println("========================")
	}
	return pids
}
