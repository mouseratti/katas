package node

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileNode struct {
	os.FileInfo
	prefix string
}

func (n *FileNode) GetChildren() []Node {
	if !n.IsDir() {
		return nil
	}
	result, err := ioutil.ReadDir(n.GetPath())
	if err != nil {
		return nil
	}
	return NewSlice(result, n.GetPath())
}

func (n *FileNode) GetPath() string {
	return fmt.Sprintf("%v/%v", n.prefix, n.Name())
}
func (n *FileNode) GetPrefixPath() string {
	return n.prefix
}

func New(i *os.FileInfo, prefix string) Node {
	f := new(FileNode)
	f.FileInfo = *i
	f.prefix = prefix
	return f
}

func NewSlice(slice []os.FileInfo, prefix string) []Node {
	length := len(slice)
	if length == 0 {
		return nil
	}
	result := make([]Node, length)

	for pos, val := range slice {
		result[pos] = New(&val, prefix)
	}
	return result
}
