package main

import (
	"aoc/helpers"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	name        string
	value       int
	parent      string
	children    []string
	accChildren int
}

func main() {
	nodes := getNodes()
	fmt.Println("Task 01: ", findRoot(nodes))
	Task02(nodes)
}

func getNodes() map[string]Node {
	data := helpers.GetLines("Data/Day07.txt")
	nodes := make(map[string]Node)
	for _, line := range data {
		split := strings.Split(line, " (")
		name := split[0]
		split = strings.Split(split[1], ")")
		value, _ := strconv.Atoi(split[0])
		children := make([]string, 0)
		if len(split[1]) > 0 {
			children = strings.Split(split[1][4:], ", ")
		}
		parent := "."
		nodes[name] = Node{name, value, parent, children, 0}
	}
	for key, value := range nodes {
		for _, child := range value.children {
			node := nodes[child]
			node.parent = key
			nodes[child] = node
		}
	}
	return nodes
}

func findRoot(nodes map[string]Node) string {
	for _, node := range nodes {
		if node.parent == "." {
			return node.name
		}
	}
	return "NotFound"
}

func Task02(nodes map[string]Node) {
	nodeList := make([]string, len(nodes))
	idx := 0
	idxInsert := 1
	currNode := findRoot(nodes)
	nodeList[0] = currNode
	for idx < len(nodes) {
		children := nodes[currNode].children
		for _, child := range children {
			nodeList[idxInsert] = child
			idxInsert++
		}
		idx++
		if idx == len(nodes) {
			idx--
			break
		}
		currNode = nodeList[idx]
	}

	for idx >= 0 {
		node := nodes[nodeList[idx]]
		acc := 0
		childrenValues := make([]int, len(node.children))
		for childIdx, child := range node.children {
			acc += nodes[child].value + nodes[child].accChildren
			childrenValues[childIdx] = nodes[child].value + nodes[child].accChildren
		}
		idxUnbal, diff := idxUnbalanced(childrenValues)
		if idxUnbal != -1 {
			fmt.Println("Task 02: ", nodes[node.children[idxUnbal]].value-diff)
			break
		}
		node.accChildren = acc
		nodes[nodeList[idx]] = node
		idx--
	}
}

func idxUnbalanced(values []int) (int, int) {
	for idx, val := range values {
		prev := idx - 1
		if prev < 0 {
			prev = len(values) - 1
		}
		next := (idx + 1) % len(values)
		if val != values[prev] && val != values[next] {
			return idx, val - values[prev]
		}
	}
	return -1, 0
}
