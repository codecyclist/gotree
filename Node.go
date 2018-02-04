package gotree

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type Node struct {
	Id       uuid.UUID
	Label    string
	Tree     *Tree `json:"-"`
	Parent   *Node `json:"-"`
	Children []*Node
	Data     interface{}
}

// NodeCriteria allows to specify a criteria for matching nodes in search queries
type NodeCriteria func(subject *Node) (found bool)

// AddChild adds a child to the current node and returns a reference to the child
func (node *Node) AddChild(id uuid.UUID, label string, data interface{}) *Node {
	newNode := node.Tree.NewNode(id, label, data)
	newNode.Parent = node
	node.Children = append(node.Children, newNode)
	return newNode
}

// AddChildren adds n children to the current node and returns a reference to it
func (node *Node) AddChildren(nodes ...*Node) (self *Node) {
	for _, newNode := range nodes {
		node.Parent = node
		node.Children = append(node.Children, newNode)
	}
	return node
}

// GetChildren returns the child node of a given node matching a criteria
func (node *Node) GetChildren(criteria NodeCriteria, maxMatches int) (results []*Node, numOfMatches int) {
	for _, newNode := range node.Children {
		if criteria(newNode) == true {
			results = append(results, newNode)
			numOfMatches++
			if maxMatches > 0 && numOfMatches == maxMatches {
				return
			}
		}
	}
	return
}

// GetChild returns the first child node of a given node matching a criteria
func (node *Node) GetChild(criteria NodeCriteria) (result *Node, found bool) {
	if matches, n := node.GetChildren(criteria, 1); n > 0 {
		result, found = matches[0], true
	}
	return
}

// GetChildByLabel returns a reference to a child of a certain Node with the title specified
func (node *Node) GetChildByLabel(label string) (needle *Node, found bool) {
	return node.GetChild(func(subject *Node) bool {
		return subject.Label == label
	})
}

// GetChildByLabel returns a reference to a child of a certain Node with the title specified
func (node *Node) GetChildByUuid(uuid uuid.UUID) (needle *Node, found bool) {
	return node.GetChild(func(subject *Node) bool {
		return subject.Id == uuid
	})
}

func (node *Node) InsertAtPath(path []string, nodes []*Node) (err error) {
	if len(path) == 0 {
		for _, newNode := range nodes {
			node.AddChildren(newNode)
		}
	} else {
		if _, exists := node.GetChildByLabel(path[0]); !exists {
			node.AddChild(uuid.New(), path[0], nil)
		}
		child, _ := node.GetChildByLabel(path[0])
		child.InsertAtPath(path[1:], nodes)
	}
	return
}

// FindByPath returns a reference to a child node relative to the current node expressed by a "child1/child2/childn" reference
func (node *Node) FindByPath(path string) (needle *Node, exists bool) {
	tokens := strings.Split(path, "/")[1:]
	return node.FindByPathTokens(tokens)
}

// FindByPathTokens returns a reference to a child node relative to the current node expressed by a slice of child node titles
func (node *Node) FindByPathTokens(tokens []string) (needle *Node, exists bool) {
	exists = false
	if next, found := node.GetChildByLabel(tokens[0]); found {
		if len(tokens) == 2 {
			needle, exists = next.GetChildByLabel(tokens[1])
		} else {
			return next.FindByPathTokens(tokens[1:])
		}
	}
	return
}

func (node *Node) String() (result string) {
	nodeJSON, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, nodeJSON, "=", "\t")
	result = out.String()
	return
}
