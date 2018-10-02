package node

import "os"

type Node interface {
	os.FileInfo
	GetChildren() []Node
	GetParents() []Node
	IsLast() bool
	GetPrefixPath() string
	SetLast()
	//FormattedSize() string
}
