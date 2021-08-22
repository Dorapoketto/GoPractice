package main

import "fmt"

func main() {
	team := map[string]float32{
		"P1": 1.98,
		"P2": 2.05,
		"P3": 1.89,
		"P4": 2.0,
		"P5": 2.11}

	var num int
	var sum float32
	for _, v := range team {
		sum += v
		num++
	}
	fmt.Println(sum / (float32(num)))

}
