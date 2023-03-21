package main

import (
	"fmt"
	"math"
)

const maxDif = 0.0000001

func Sqrt(x float64) float64 {
	z := 1.0
	for dif := math.Inf(1); maxDif <= math.Abs(dif); dif = z*z - x {
		z -= dif / (z * 2)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
