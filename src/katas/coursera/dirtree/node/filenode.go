package node

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileNode struct {
	os.FileInfo
	prefix string
	isLast bool
	parent Node
}

func (n *FileNode) SetLast() {
	n.isLast = true
}

func (n *FileNode) IsLast() bool {
	return n.isLast
}

func (n *FileNode) GetParents() []Node {
	if n.parent == nil {
		return nil
	}
	result := make([]Node, 1)
	result[0] = n.parent
	return append(result, n.parent.GetParents()...)
}

func (n *FileNode) GetChildren() []Node {
	if !n.IsDir() {
		return nil
	}
	result, err := ioutil.ReadDir(n.GetPath())
	if err != nil {
		return nil
	}
	return NewSlice(result, n.GetPath(), n)
}

func (n *FileNode) GetPath() string {
	return fmt.Sprintf("%v/%v", n.prefix, n.Name())
}
func (n *FileNode) GetPrefixPath() string {
	return n.prefix
}

func New(i *os.FileInfo, prefix string, isLast bool, parent Node) Node {
	f := new(FileNode)
	f.FileInfo = *i
	f.prefix = prefix
	f.isLast = isLast
	f.parent = parent
	return f
}

func NewSlice(slice []os.FileInfo, prefix string, parent Node) []Node {
	length := len(slice)

	if length == 0 {
		return nil
	}
	maxIndex := length - 1
	result := make([]Node, length)

	for pos, val := range slice {
		isLast := maxIndex == pos
		result[pos] = New(&val, prefix, isLast, parent)
	}
	return result
}
