// Package edit defines the edit-script types shared by the diff pipeline
// steps and by the renderers that consume their output.
package edit

import (
	"fmt"
)

// EditNode is one entry in an edit script: a single line and the
// operation that transforms the old file into the new one at that point.
// The zero value is an empty, unchanged line.
type EditNode struct {
	// Content is the text of the line, without its line terminator.
	Content string
	// Op says whether the line is unchanged, added, or removed.
	Op EditOp
}

// EditOp is the kind of change applied to a line.
type EditOp int

const (
	// NoChange marks a line present in both files. It is the zero value.
	NoChange EditOp = iota
	// Add marks a line present only in the new file.
	Add
	// Remove marks a line present only in the old file.
	Remove
)

// String returns the constant's name, for debugging and test output.
// Display symbols such as "+" and "-" belong to the render package.
func (op EditOp) String() string {
	switch op {
	case NoChange:
		return "NoChange"
	case Add:
		return "Add"
	case Remove:
		return "Remove"
	default:
		return fmt.Sprintf("Unknown LineOperation(%d)", int(op))
	}
}
