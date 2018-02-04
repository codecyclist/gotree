# Tree Data Container for Go
Gotree is a set of convenience functions to build an in-memory tree structure in Go to organize data of any sort.

## Examples

### Creating a Tree for managing machine settings
```
// Create a new Tree
tree := NewTree()

// Create the nodes /Machine/Heating/Zones/Zone1/SetTemperature and ActualTemperature
// (nil in lieu of the actual data)
tree.Root.InsertAtPath(
		[]string{"Machine", "Heating", "Zones", "Zone1"},
		[]*Node{
			tree.NewNode(uuid.New(), "SetTemperature", Opc.get("V27t1")),
			tree.NewNode(uuid.New(), "ActualTemperature", Opc.get("V2t87")),
		}

```

### Searching for Nodes
```
// Is there a /robots/marvin/quotes in our tree?
needle, found := tree.Root.FindByPathTokens([]string{"robots", "marvin", "quotes"})
```

Please check the demo and the unittests for more examples.