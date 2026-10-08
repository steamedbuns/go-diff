package edit

import (
	"testing"
)

func TestLineOperationString(t *testing.T) {
	tests := []struct {
		name string
		in   Op
		want string
	}{
		{name: "NoChange", in: NoChange, want: "NoChange"},
		{name: "Add", in: AddLine, want: "AddLine"},
		{name: "Remove", in: RemoveLine, want: "RemoveLine"},
		{name: "Default", in: Op(12), want: "Unknown LineOperation(12)"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.in.String()
			if got != tc.want {
				t.Errorf("LineOperation.String(%v) = %v, want %v", int(tc.in), got, tc.want)
			}
		})
	}
}

func TestNodeHelperFunction(t *testing.T) {
	tests := []struct {
		name     string
		callFunc func(s string) Node
		want     Node
	}{
		{name: "Keep", callFunc: Keep, want: Node{Content: "Keep", Op: NoChange}},
		{name: "Add", callFunc: Add, want: Node{Content: "Add", Op: AddLine}},
		{name: "Remove", callFunc: Remove, want: Node{Content: "Remove", Op: RemoveLine}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.callFunc(tc.name)
			if got != tc.want {
				t.Errorf("%v(%v) = %v, want %v", tc.name, tc.name, got, tc.want)
			}
		})
	}
}
