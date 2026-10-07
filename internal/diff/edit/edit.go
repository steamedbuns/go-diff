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
		return "LineAdded"
	case Remove:
		return "LineRemoved"
	default:
		return fmt.Sprintf("Unknown LineOperation(%d)", int(op))
	}
}
