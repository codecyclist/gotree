package gotree

import ("testing")
import "github.com/google/uuid"

func TestCreate(t *testing.T) {
	tree := NewTree()
	t.Logf("tree=", tree)
}

func TestFindByPath(t *testing.T) {
	tree := NewTree()

	tree.Root.AddChild(uuid.New(), "robots", nil).
	AddChild(uuid.New(), "marvin", nil).
	AddChild(uuid.New(), "quotes", []string {"Robot1"})

	node, found := tree.FindByPath("/robots/marvin/quotes")

	if !found || node.Data.([]string)[0]!="Robot1" {
		t.Fail()
	}

}