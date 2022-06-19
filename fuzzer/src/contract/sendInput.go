package contract

import (
	"Fuzzer/src/jsonData"
	"Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

func SendInput(inputNew pidType.INPUTDATA, method string, ContractAddress string, abi abi.ABI, CodeId int) (pidType.COST, []int64, []string) {
	inputPacked := tools.InputEncode(inputNew, method, abi)
	dataString := hex.EncodeToString(inputPacked)

	requestJSON, err := json.Marshal(jsonData.TransactionReq{CodeId: CodeId, InputData: dataString})
	if err != nil {
		panic(err)
	}

	res, err := http.Post(
		"http://192.168.159.136:5000/api/job/test/",
		"application/json;charset=UTF-8",
		strings.NewReader(string(requestJSON)),
	)
	if err != nil {
		panic(err)
	}

	defer func() { _ = res.Body.Close() }()

	responseJSON, _ := ioutil.ReadAll(res.Body)

	var transaction jsonData.TransactionRes
	err = json.Unmarshal(responseJSON, &transaction)

	if err != nil {
		panic(err)
	}

	var ResultData jsonData.Result

	err = json.Unmarshal([]byte(transaction.Response), &ResultData)

	if err != nil {
		panic(err)
	}

	var pcInt []int64
	var pathInt []int64
	var costValue []*big.Int
	for _, pcstr := range ResultData.Pc {
		i, _ := strconv.ParseInt(pcstr, 0, 64)
		pcInt = append(pcInt, i)
	}
	for _, coststr := range ResultData.Cost {
		var i = big.NewInt(0)
		i, _ = i.SetString(coststr[2:], 16)
		costValue = append(costValue, i)
	}
	for _, pathstr := range ResultData.Path {
		i, _ := strconv.ParseInt(pathstr, 0, 64)
		pathInt = append(pathInt, i)
	}

	cost := tools.Cost2map(pcInt, costValue)

	return cost, pathInt, ResultData.Status
}
