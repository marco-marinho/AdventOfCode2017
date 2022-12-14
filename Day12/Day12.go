package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := helpers.GetLines("Data/Day12.txt")
	nodes := make(map[int][]int)
	for _, line := range data {
		node, _ := strconv.Atoi(strings.Split(line, " <")[0])
		connectionsStr := strings.Split(strings.Split(line, "> ")[1], ", ")
		connections := make([]int, len(connectionsStr)+1)
		for idx, nStr := range connectionsStr {
			connections[idx+1], _ = strconv.Atoi(nStr)
		}
		connections[0] = node
		nodes[node] = connections
	}
	connectedToZero := getGroupConnected(nodes, 0)
	fmt.Println("Task 01:", len(connectedToZero))
	for key := range connectedToZero {
		delete(nodes, key)
	}
	groups := 1
	for len(nodes) > 0 {
		root := getSomeKey(nodes)
		connectedToRoot := getGroupConnected(nodes, root)
		for key := range connectedToRoot {
			delete(nodes, key)
		}
		groups += 1
	}
	fmt.Println("Task 02:", groups)
}

func getGroupConnected(nodes map[int][]int, root int) map[int]bool {
	connectedToRoot := make(map[int]bool)
	for _, element := range nodes[root] {
		connectedToRoot[element] = true
	}
	sizeConnectedStart := len(connectedToRoot)
	sizeConnectedStop := 0
	for sizeConnectedStart != sizeConnectedStop {
		sizeConnectedStart = len(connectedToRoot)
		for key, value := range nodes {
			_, ok := connectedToRoot[key]
			if ok {
				for _, connections := range value {
					connectedToRoot[connections] = true
				}
			}
		}
		sizeConnectedStop = len(connectedToRoot)
	}
	return connectedToRoot
}

func getSomeKey(m map[int][]int) int {
	for k := range m {
		return k
	}
	return 0
}
