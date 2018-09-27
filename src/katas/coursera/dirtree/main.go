package main

import (
	"fmt"
	"os"
)

import "io"
import "io/ioutil"
import "katas/coursera/dirtree/node"

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	result, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	firstLevelnodes := node.NewSlice(result, path)
	var stringSlice []string
	stringSlice = node.Print(firstLevelnodes, printFiles)
	fmt.Println(stringSlice)
	return nil
}
