package jsonData

type DeployReq struct {
	Name     string      `json:"name"`
	AbiArray interface{} `json:"abi"`
	Bytecode string      `json:"bytecodes"`
	Data     string      `json:"data"`
}

type DeployRes struct {
	CodeId   int         `json:"id"`
	Name     string      `json:"name"`
	AbiArray interface{} `json:"abi"`
	Address  string      `json:"address"`
}

type TransactionReq struct {
	CodeId    int    `json:"id"`
	InputData string `json:"input"`
}

type TransactionRes struct {
	CodeId   int    `json:"id"`
	Response string `json:"response"`
}

type Result struct {
	Pc     []string `json:"pc"`
	Path   []string `json:"path"`
	Cost   []string `json:"cost"`
	Status []string `json:"status"`
}
