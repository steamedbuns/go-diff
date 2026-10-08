package normalize

import "github.com/steamedbuns/go-diff/internal/diff/edit"

func EditNodeGrouping(nodes []edit.Node) []edit.Node {
	addNodes := make([]edit.Node, len(nodes))[:0]
	normalized := make([]edit.Node, len(nodes))[:0]
	changeChunk := false
	for _, node := range nodes {
		switch node.Op {
		case edit.NoChange:
			if changeChunk {
				changeChunk = !changeChunk
				normalized = append(normalized, addNodes...)
				addNodes = addNodes[:0]
			}
			normalized = append(normalized, node)
		case edit.AddLine:
			if !changeChunk {
				changeChunk = !changeChunk
			}
			addNodes = append(addNodes, node)
		case edit.RemoveLine:
			if !changeChunk {
				changeChunk = !changeChunk
			}
			normalized = append(normalized, node)
		}
	}
	if len(addNodes) > 0 {
		normalized = append(normalized, addNodes...)
	}
	return normalized
}
