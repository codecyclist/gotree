package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/codecyclist/gotree"
	"github.com/google/uuid"
	"os"
)

type OpcValue struct {
	WellknownId     string
	EngineeringUnit string
	Quality         Quality
	Payload         string
}

type Quality bool

const (
	Good Quality = true
	Bad  Quality = false
)

func main() {
	tree := gotree.NewTree()

	pressures := tree.Root.AddChild(
		uuid.New(),
		"pressures",
		nil,
	)

	pressures.AddChild(
		uuid.New(),
		"P42",
		OpcValue{
			EngineeringUnit: "bar",
			WellknownId:     "P42",
			Quality:         Good,
			Payload:         "uint32:23.23",
		})

	treeJson, _ := json.Marshal(tree)
	var out bytes.Buffer
	json.Indent(&out, treeJson, "=", "\t")
	out.WriteTo(os.Stdout)

	if node, exists := tree.Root.FindByPath("/pressures/P42"); exists {
		fmt.Println("Found Node:", node, " value=", node.Data.(OpcValue).Payload)
	}

	if node, exists := tree.Root.FindByPath("/not/existing/at/all"); exists {
		fmt.Println("Found Node:", node)
	}

}
