package main

import (
	"fmt"
	"math"
)

func main() {
	infinite := math.Inf(1)
	graph := map[string]map[string]int{
		"Start": map[string]int{
			"A": 5,
			"B": 2,
		},
		"A": map[string]int{
			"C": 4,
			"D": 2,
		},
		"B": map[string]int{
			"A": 8,
			"D": 7,
		},
		"C": map[string]int{
			"D":   6,
			"End": 3,
		},
		"D": map[string]int{
			"End": 1,
		},
		"End": map[string]int{},
	}

	costs := map[string]float64{
		"A":   5,
		"B":   2,
		"End": infinite,
	}

	parents := map[string]string{
		"A":   "start",
		"B":   "start",
		"End": "",
	}

	fmt.Print(infinite, "\n")
	fmt.Print(graph, "\n")
	fmt.Print(costs, "\n")
	fmt.Print(parents, "\n")
}
