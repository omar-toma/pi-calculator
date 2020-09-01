package main

import (
	"math/big"
	"fmt"
)

var (
	digits  = 10000
)
var pi *big.Float

func main() {
	pi, _ = CalcPi(float64(digits))
	fmt.Printf("%1.[1]*[2]f", digits, pi)
}
