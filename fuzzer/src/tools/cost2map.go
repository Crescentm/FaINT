package tools

import (
	"Fuzzer/src/pidType"
	"math/big"
)

func Cost2map(pc []int64, costValue []*big.Int) pidType.COST {
	cost := make(pidType.COST)
	for i := 0; i < len(pc); i = i + 2 {
		cost[pc[i]] = append(cost[pc[i]], costValue[i])
		cost[pc[i]] = append(cost[pc[i]], costValue[i+1])
	}
	return cost
}
