package diff

import (
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

var lcsTests = []struct {
	name   string
	inputA []string
	inputB []string
	want   []edit.Node
}{
	{
		name:   "Placeholder",
		inputA: []string{"a", "b", "c"},
		inputB: []string{},
		want: []edit.Node{
			edit.Remove("a"),
			edit.Remove("b"),
			edit.Remove("c"),
		},
	},
}

func TestLCS(t *testing.T) {
	for _, tc := range lcsTests {
		t.Run(tc.name, func(t *testing.T) {
			got := LCS(tc.inputA, tc.inputB)
			if len(got) != len(tc.want) {
				t.Errorf("LCS(%v,%v) = %v, want %v", tc.inputA, tc.inputB, got, tc.want)
			}
		})
	}
}
