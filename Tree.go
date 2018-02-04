package gotree

import (
	"strings"

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

// NewNode creates a new detached node
func (t *Tree) NewNode(id uuid.UUID, label string, data interface{}) *Node {
	// Sanitize Label and gracefully replace with id if empty
	label = strings.Replace(label, "/", "", -1)
	if len(label) == 0 {
		label = id.String()
	}

	newNode := &Node{
		Tree:   t,
		Id:     id,
		Label:  label,
		Parent: nil,
		Data:   data,
	}
	return newNode
}

// FindByPath returns a Node by a path combined from the node titles
func (t *Tree) FindByPath(path string) (needle *Node, exists bool) {
	return t.Root.FindByPath(path)
}

func (t *Tree) String() (result string) {
	return t.Root.String()
}
