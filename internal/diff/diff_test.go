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
	{"empty", []string{}, nil},
	{"same data", []string{"a", "a", "a"}, []int{0, 0, 0}},
	{"different data", []string{"a", "b", "c"}, []int{0, 1, 2}},
	{"data shuffle", []string{"a", "b", "a", "c"}, []int{0, 1, 0, 2}},
}

func TestTokenize(t *testing.T) {
	for _, tc := range tokenizeTests {
		t.Run(tc.name, func(t *testing.T) {
			tokenMap := newTokenMap()
			got := tokenMap.tokenize(tc.input)
			if !slices.Equal(got, tc.want) {
				t.Errorf("Tokenize(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

var lcsTests = []struct {
	name   string
	inputA []int
	inputB []int
	want   []edit.Node
}{
	{"both empty", []int{}, []int{}, []edit.Node{}},
	{"inputA empty", []int{}, []int{1, 2, 3, 4}, []edit.Node{edit.Add(1), edit.Add(2), edit.Add(3), edit.Add(4)}},
	{"inputB empty", []int{1, 2, 3, 4}, []int{}, []edit.Node{edit.Remove(1), edit.Remove(2), edit.Remove(3), edit.Remove(4)}},
	{"all equal", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []edit.Node{edit.Keep(1), edit.Keep(2), edit.Keep(3), edit.Keep(4)}},
	{"start equal", []int{1, 2, 5, 5}, []int{1, 2, 6, 6}, []edit.Node{edit.Keep(1), edit.Keep(2), edit.Remove(5), edit.Remove(5), edit.Add(6), edit.Add(6)}},
	{"end equal", []int{5, 5, 1, 2}, []int{6, 6, 1, 2}, []edit.Node{edit.Remove(5), edit.Remove(5), edit.Add(6), edit.Add(6), edit.Keep(1), edit.Keep(2)}},
	{"middle equal", []int{5, 1, 2, 5}, []int{6, 1, 2, 6}, []edit.Node{edit.Remove(5), edit.Add(6), edit.Keep(1), edit.Keep(2), edit.Remove(5), edit.Add(6)}},
	{"multiple equal", []int{1, 2, 5, 5, 5, 6, 7, 8}, []int{1, 2, 4, 4, 6, 7, 8}, []edit.Node{edit.Keep(1), edit.Keep(2), edit.Remove(5), edit.Remove(5), edit.Remove(5), edit.Add(4), edit.Add(4), edit.Keep(6), edit.Keep(7), edit.Keep(8)}},
}

func TestLcs(t *testing.T) {
	for _, tc := range lcsTests {
		t.Run(tc.name, func(t *testing.T) {
			got := lcs(tc.inputA, tc.inputB)
			if !slices.Equal(got, tc.want) {
				t.Errorf("lcs(%v, %v) =\ngot  %v\nwant %v", tc.inputA, tc.inputB, got, tc.want)
			}
		})
	}
}
