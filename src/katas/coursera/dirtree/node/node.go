package node

import "os"

type Node interface {
	os.FileInfo
	GetChildren() []Node
	GetPath() string
	GetPrefixPath() string
	//FormattedSize() string
}
