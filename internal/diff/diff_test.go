package diff

import (
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		name   string
		inputA []string
		inputB []string
		want   []edit.Node
	}{
		{name: "Placeholder", inputA: []string{"a", "b", "c"}, inputB: []string{}, want: []edit.Node{
			{Content: "a", Op: edit.Remove},
			{Content: "b", Op: edit.Remove},
			{Content: "c", Op: edit.Remove}}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LCS(tc.inputA, tc.inputB)
			if len(got) != len(tc.want) {
				t.Errorf("LCS(%v,%v) = %v, want %v", tc.inputA, tc.inputB, got, tc.want)
			}
		})
	}
}
