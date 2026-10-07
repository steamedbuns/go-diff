package edit

import (
	"fmt"
)

type LineOperation int

const (
	NoChange LineOperation = iota
	Add
	Remove
)

func (op LineOperation) String() string {
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
