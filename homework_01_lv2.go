package main

import (
	"fmt"
	"math"
)

func calculateCircleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
func main() {
	r := 1.0
	a := calculateCircleArea(r)
	fmt.Println("面积：", a)
}
