package main

import "math/big"

func addBinary(a string, b string) string {
	intA := toInt(a)
	intB := toInt(b)
	sum := new(big.Int).Add(&intA, &intB)
	return toString(*sum)
}

func toInt(a string) big.Int {
	intA, _ := new(big.Int).SetString(a, 2)
	return *intA
}

func toString(a big.Int) string {
	return a.Text(2)
}

func main() {
}
