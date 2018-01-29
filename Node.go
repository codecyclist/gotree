package gotree

import (
	"github.com/google/uuid"
	"strings"
	"bytes"
	"encoding/json"
	"os"
)

type Node struct {
	Id       uuid.UUID
	Label    string
	Parent   *Node `json:"-"`
	Children []*Node
	Data     interface{}
}

func (node *Node) AddChild(id uuid.UUID, label string, data interface{}) *Node {

	newNode := &Node{
		Id:     id,
		Label:  label,
		Parent: node,
		Data:   data,
	}
	node.Children = append(node.Children, newNode)

	return newNode
}

func (node *Node) FindByPath(path string) (needle Node, exists bool) {
	exists = false
	tokens := strings.Split(path, "/")[1:]

	for _, child := range node.Children {
		if child.Label == tokens[0] {
			if len(tokens) == 1 {
				needle = *child
				exists = true
				return
			} else {
				return child.FindByPath(strings.Join(tokens, "/"))
			}
		}
	}
	return
}


func(node *Node) ToJson() {
	nodeJson, _ := json.Marshal(node)
	var out bytes.Buffer
	json.Indent(&out, nodeJson, "=", "\t")	
	out.WriteTo(os.Stdout)
}
