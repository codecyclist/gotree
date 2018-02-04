package gotree

import (
	"testing"
)
import "github.com/google/uuid"

func TestCreate(t *testing.T) {
	tree := NewTree()
	t.Logf("tree=%v", tree.String())
}

func TestCreateWithMultipleChildrenInsert(t *testing.T) {
	tree := NewTree()
	tree.Root.AddChild(uuid.New(), "robots", nil).
		AddChildren(
			tree.NewNode(uuid.New(), "Vincent", nil),
			tree.NewNode(uuid.New(), "Marvin", nil),
		)

	tree.Root.AddChild(uuid.New(), "movies", nil).
		AddChildren(
			tree.NewNode(uuid.New(), "Hitchhikers Guide to Galaxy", nil),
			tree.NewNode(uuid.New(), "Black Hole", nil),
		)

	t.Logf("tree=%v", tree.String())
}

func TestFindByPath(t *testing.T) {
	tree := NewTree()

	tree.Root.AddChild(uuid.New(), "robots", nil).
	AddChild(uuid.New(), "marvin", nil).
		AddChild(uuid.New(), "quotes", []string{"Robot1"})

	node, found := tree.FindByPath("/robots/marvin/quotes")

	if !found || node.Data.([]string)[0] != "Robot1" {
		t.Fail()
	}
}

func TestFindByPathTokens(t *testing.T) {
	tree := NewTree()

	tree.Root.AddChild(uuid.New(), "robots", nil).
		AddChild(uuid.New(), "marvin", nil).
		AddChild(uuid.New(), "quotes", "MARVIN")

	needle, found := tree.Root.FindByPathTokens([]string{"robots", "marvin", "quotes"})

	if !found || needle.Data.(string) != "MARVIN" {
		t.Fail()
	} else {
		t.Logf("found=%v", needle)
	}
}

func TestFindChildByLabelWithExistingChild(t *testing.T) {
	tree := NewTree()

	tree.Root.AddChild(uuid.New(), "robots", nil).
		AddChild(uuid.New(), "marvin", nil).
		AddChild(uuid.New(), "quotes", []string{"Robot1"})

	if needle, exists := tree.Root.GetChildByLabel("robots"); exists {
		t.Logf("found=%v", needle)
	} else {
		t.Fail()
	}
}

func TestFindChildByUuidlWithExistingChild(t *testing.T) {
	tree := NewTree()
	existingId := uuid.New()
	tree.Root.AddChild(existingId, "robots", nil)

	if needle, exists := tree.Root.GetChildByUuid(existingId); exists {
		t.Logf("found=%v", needle)
	} else {
		t.Fail()
	}
}

func TestFindChildByUuidlWithNonExistingChild(t *testing.T) {
	tree := NewTree()
	existingId := uuid.New()
	tree.Root.AddChild(existingId, "robots", nil)

	if _, exists := tree.Root.GetChildByUuid(uuid.New()); !exists {
		t.Logf("not found as expected")
	} else {
		t.Fail()
	}
}

func TestInsertAtPathCreateNewNodeInEmptyTree(t *testing.T) {
	tree := NewTree()

	if error := tree.Root.InsertAtPath(
		[]string{"Machine", "Heating", "Zones", "Zone1"},
		[]*Node{
			tree.NewNode(uuid.New(), "SetTemperature", nil),
			tree.NewNode(uuid.New(), "ActualTemperature", nil),
		}); error != nil {
		t.Fail()
	}

	if error := tree.Root.InsertAtPath(
		[]string{"Machine", "Heating", "Zones", "Zone2"},
		[]*Node{
			tree.NewNode(uuid.New(), "SetTemperature", nil),
			tree.NewNode(uuid.New(), "ActualTemperature", nil),
		}); error != nil {
		t.Fail()
	}

	if _, exists := tree.FindByPath("/Machine/Heating/Zones/Zone2/SetTemperature"); !exists {
		t.Fatalf("Expected node not found!")
	}

	if _, exists := tree.FindByPath("/Machine/Heating/Zones"); !exists {
		t.Fatalf("Node on the way not found!")
	}

	t.Logf("tree=%v", tree.String())
}

func TestFindByPathShouldReturnFalseOnEmptyTree(t *testing.T) {
	tree := NewTree()

	if _, found := tree.FindByPath("/robots/marvin/quotes"); found {
		t.Fail()
	}

}