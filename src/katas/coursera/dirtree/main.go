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
	firstLevelnodes := node.NewSlice(result, path, nil)
	buff := node.Print(firstLevelnodes, printFiles)
	for i := range buff {
		fmt.Println(i)
	}
	return nil
}
