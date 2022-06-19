package fuzzer

import (
	"Fuzzer/src/contract"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"time"
)

var (
	AbiObject       abi.ABI
	ABIFile         *string
	SeedFile        *string
	BinaryFile      *string
	ContractAddress string
	CodeId          int
)

func Init() {
	ABIFile = flag.String("abi", "Test.abi", "abi file path")
	SeedFile = flag.String("seed", "seed.json", "seed file path")
	BinaryFile = flag.String("binary", "Test.bin-runtime", "binary file path")
}

func Fuzz() {
	start := time.Now()
	flag.Parse()

	ABIFileStream, err := os.Open(*ABIFile)
	if err != nil {
		panic("No such abi file")
	}

	AbiObject, err = abi.JSON(ABIFileStream)
	if err != nil {
		panic("Can not parse abi file")
	}

	err = ABIFileStream.Close()
	if err != nil {
		panic(err)
	}

	res, err := contract.Deploy(*BinaryFile, *ABIFile, AbiObject)
	if err != nil {
		panic(err)
	}

	ContractAddress = res.Address
	CodeId = res.CodeId

	for fuctionName := range AbiObject.Methods {
		if len(AbiObject.Methods[fuctionName].Inputs) < 1 {
			continue
		}
		FuctionFuzz(fuctionName, SeedFile)
	}
	end := time.Now()
	sub := end.Sub(start)
	fmt.Println("===========time=============")
	fmt.Printf("time: %4fs\n", sub.Seconds())
	fmt.Println("===========time=============")
}
