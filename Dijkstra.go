package main

import (
	"fmt"
	"math"
)

func getKeys(hash map[string]int) []string {
	keys := make([]string, len(hash))
	i := 0
	for key := range hash {
		keys[1] = key
		i++
	}

	return keys
}

func indexOf(arr []string, item string) int {
	for index, value := range arr {
		if value == item {
			return index
		}
	}

	return -1
}

func dijkstra(graph map[string]map[string]int, startingNode string, endNode string) {
	infinite := math.Inf(1)
	costs := map[string]int{}
	parents := map[string]string{}
	var tracked []string

	for key, value := range graph[startingNode] {
		costs[key] = value
		parents[key] = startingNode
	}
	costs["End"] = int(infinite)
	parents["End"] = ""

	tracked = append(tracked, startingNode)

	fmt.Print(graph, "\n")
	fmt.Print(costs, "\n")
	fmt.Print(parents, "\n")
	fmt.Print(tracked, "\n")

	fmt.Print(indexOf(tracked, startingNode), "\n")
	fmt.Print(indexOf(tracked, endNode), "\n")
}

func main() {
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

	dijkstra(graph, "Start", "End")
}
