package main

import (
	"fmt"
	"math"
)

func getKeys(hash map[string]float64) []string {
	keys := make([]string, 0, len(hash))
	for key := range hash {
		keys = append(keys, key)
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

func reverseArr(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}

func findCheapestNode(hash map[string]float64, track []string) string {
	lowestNode := ""
	lowestCost := math.Inf(0)

	for node := range hash {
		nodeCost := hash[node]
		if (nodeCost < lowestCost) && indexOf(track, node) == -1 {
			lowestCost = nodeCost
			lowestNode = node
		}
	}

	return lowestNode
}

func setInitialCostsParents(hash map[string]float64, startingNode string, endNode string) (map[string]float64, map[string]string) {
	costs := map[string]float64{}
	parents := map[string]string{}
	for key, value := range hash {
		costs[key] = value
		parents[key] = startingNode
	}
	costs[endNode] = math.Inf(0)
	parents[endNode] = ""

	return costs, parents
}

func buildFinalRoute(hash map[string]string, startingNode string, activeNode string) []string {
	route := make([]string, 0, len(hash))

	for activeNode != startingNode {
		route = append(route, activeNode)
		activeNode = hash[activeNode]
	}
	route = append(route, startingNode)
	route = reverseArr(route)

	return route
}

func dijkstra(graph map[string]map[string]float64, startingNode string, endNode string) []string {
	costs, parents := setInitialCostsParents(graph[startingNode], startingNode, endNode)
	tracked := make([]string, 0, len(graph))
	tracked = append(tracked, startingNode)
	activeNode := findCheapestNode(costs, tracked)

	for activeNode != endNode {
		neighbors := graph[activeNode]
		activeNodeCost := costs[activeNode]

		for key, neighborCost := range neighbors {
			newCost := activeNodeCost + neighborCost
			if costs[key] == 0 {
				costs[key] = math.Inf(0)
			}
			if newCost < costs[key] {
				costs[key] = newCost
				parents[key] = activeNode
			}
		}
		tracked = append(tracked, activeNode)
		activeNode = findCheapestNode(costs, tracked)
	}

	route := buildFinalRoute(parents, startingNode, activeNode)

	return route
}

func main() {
	graph := map[string]map[string]float64{
		"Start": map[string]float64{
			"A": 5,
			"B": 2,
		},
		"A": map[string]float64{
			"C": 4,
			"D": 20,
		},
		"B": map[string]float64{
			"A": 1,
			"D": 7,
		},
		"C": map[string]float64{
			"D":   10,
			"End": 3,
		},
		"D": map[string]float64{
			"End": 1,
		},
		"End": map[string]float64{},
	}

	route := dijkstra(graph, "Start", "End")

	fmt.Print(route)
}
