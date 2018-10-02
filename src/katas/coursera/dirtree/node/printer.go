package node

import (
	"fmt"
	"strings"
)

func Print(nodes []Node, verbose bool) chan string {
	result := make(chan string)
	nodesAsStringSlice(nodes, verbose, result)
	return result
}

func nodesAsStringSlice(nodes []Node, verbose bool, buffer chan string) {
	lastElementIndex := len(nodes) - 1
	for pos, node := range nodes {
		isLast := pos == lastElementIndex
		if isLast {
			nodes[pos].SetLast()
		}
		nodeAsString(node, verbose, buffer)
		nodesAsStringSlice(node.GetChildren(), verbose, buffer)
	}
}

func nodeAsString(node Node, verbose bool, buffer chan string) {
	// do not draw files
	if !verbose && !node.IsDir() {
		return
	}
	nodeStringbuffer := make(chan string)
	switch verbose {
	case true:
		nodeStringbuffer <- fmt.Sprintf(fullTemplate, node.Name(), node.Size())

	case false:
		nodeStringbuffer <- fmt.Sprintf(shortTemplate, node.Name())
	}
	if node.IsLast() {
		nodeStringbuffer <- lastStringStart
	} else {
		nodeStringbuffer <- normalStringStart
	}

	for _, parent := range node.GetParents() {
		if parent.IsLast() {
			nodeStringbuffer <- emptySeparator
		} else {
			nodeStringbuffer <- noEmptySeparator
		}
	}
	//close(nodeStringbuffer)
	buffer <- strings.Join(func() []string {
		length := len(nodeStringbuffer)
		result := make([]string, length)
		for pos := range result {
			result[pos] = <-nodeStringbuffer
		}
		return result
	}(), "")
}
