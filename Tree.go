package gotree

import (
	"github.com/google/uuid"
)

type Tree struct {
	Root Node
}

// NewTree creates a new Tree with a Root Node. The root node gets a random uuid.
func NewTree() (tree Tree) {
	tree.Root = Node{
		Label: "/",
		Id:    uuid.New(),
	}
	return
}

func (t Tree) FindByPath(path string) (needle Node, exists bool) {
	return t.Root.FindByPath(path)
}
