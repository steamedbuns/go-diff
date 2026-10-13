package diff

import (
	"slices"
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

var tokenizeTests = []struct {
	name  string
	input []string
	want  []int
}{
	{"empty", []string{}, []int{}},
	{"same data", []string{"a", "a", "a"}, []int{0, 0, 0}},
	{"different data", []string{"a", "b", "c"}, []int{0, 1, 2}},
	{"data shuffle", []string{"a", "b", "a", "c"}, []int{0, 1, 0, 2}},
}

func TestTokenize(t *testing.T) {
	for _, tc := range tokenizeTests {
		t.Run(tc.name, func(t *testing.T) {
			tokenMap := NewTokenMap()
			got := tokenMap.Tokenize(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("Tokenize(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

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
