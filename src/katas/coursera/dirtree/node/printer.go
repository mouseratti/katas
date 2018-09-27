package node

import (
	"fmt"
	"strings"
)

func Print(nodes []Node, verbose bool) []string {
	var result []string
	var template = shortTemplate
	nestingLevel := 0
	if verbose {
		template = fullTemplate
	}
	result = nodesAsStringSlice(nodes, &template, nestingLevel)
	return result
}

func nodesAsStringSlice(nodes []Node, template *string, nestingLevel int) []string {
	var result []string
	lastElementIndex := len(nodes) - 1
	for pos, node := range nodes {
		isLast := pos == lastElementIndex
		result = append(result, nodeAsString(node, template, nestingLevel, isLast))
		result = append(result, nodesAsStringSlice(node.GetChildren(), template, nestingLevel+1)...)
	}
	return result
}

func nodeAsString(node Node, template *string, nestingLevel int, isLast bool) string {
	var result string
	separator := formatSeparartor(nestingLevel)
	stringStart := getStringStart(isLast)
	switch *template {
	//case fullTemplate:
	default:
		result = fmt.Sprintf(shortTemplate, separator, stringStart, node.Name())
	}
	return result
}

func formatSeparartor(nestingLevel int) string {
	var result string = ""
	switch nestingLevel {
	case 0:
		return result
	default:
		slice := make([]string, nestingLevel)
		for pos := range slice {
			slice[pos] = separator
		}
		result = strings.Join(slice, "")
	}
	return result
}

func getStringStart(isLast bool) string {
	switch isLast {
	case true:
		return lastStringStart
	default:
		return normalStringStart
	}
}
