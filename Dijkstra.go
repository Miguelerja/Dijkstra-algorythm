package main

import (
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

func dijkstra(graph map[string]map[string]float64, startingNode string, endNode string) {
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
}

func main() {
	graph := map[string]map[string]float64{
		"Start": map[string]float64{
			"A": 5,
			"B": 2,
		},
		"A": map[string]float64{
			"C": 4,
			"D": 2,
		},
		"B": map[string]float64{
			"A": 8,
			"D": 7,
		},
		"C": map[string]float64{
			"D":   6,
			"End": 3,
		},
		"D": map[string]float64{
			"End": 1,
		},
		"End": map[string]float64{},
	}

	dijkstra(graph, "Start", "End")
}
