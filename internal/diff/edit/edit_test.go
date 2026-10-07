package edit

import (
	"testing"
)

func TestLineOperationString(t *testing.T) {
	tests := []struct {
		name string
		in   LineOperation
		want string
	}{
		{name: "NoChange", in: NoChange, want: "NoChange"},
		{name: "Add", in: Add, want: "LineAdded"},
		{name: "Remove", in: Remove, want: "LineRemoved"},
		{name: "Default", in: LineOperation(12), want: "Unknown LineOperation(12)"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.in.String()
			if got != tc.want {
				t.Errorf("LineOperation.String(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
