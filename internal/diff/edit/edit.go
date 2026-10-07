// Package edit defines the edit-script types shared by the diff pipeline
// steps and by the renderers that consume their output.
package edit

import (
	"fmt"
)

// Node is one entry in an edit script: a single line and the
// operation that transforms the old file into the new one at that point.
// The zero value is an empty, unchanged line.
type Node struct {
	// Content is the text of the line, without its line terminator.
	Content string
	// Op says whether the line is unchanged, added, or removed.
	Op Op
}

// Op is the kind of change applied to a line.
type Op int

const (
	// NoChange marks a line present in both files. It is the zero value.
	NoChange Op = iota
	// Add marks a line present only in the new file.
	AddLine
	// Remove marks a line present only in the old file.
	RemoveLine
)

// String returns the constant's name, for debugging and test output.
// Display symbols such as "+" and "-" belong to the render package.
func (op Op) String() string {
	switch op {
	case NoChange:
		return "NoChange"
	case AddLine:
		return "AddLine"
	case RemoveLine:
		return "RemoveLine"
	default:
		return fmt.Sprintf("Unknown LineOperation(%d)", int(op))
	}
}

func Keep(s string) Node {
	return Node{Content: s, Op: NoChange}
}

func Add(s string) Node {
	return Node{Content: s, Op: AddLine}
}

func Remove(s string) Node {
	return Node{Content: s, Op: RemoveLine}
}
