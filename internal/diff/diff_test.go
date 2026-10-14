package diff

import (
	"slices"
	"testing"
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
	inputA []int
	inputB []int
	want   []int
}{
	{"both empty", []int{}, []int{}, []int{}},
	{"inputA empty", []int{}, []int{1, 2, 3, 4}, []int{}},
	{"inputB empty", []int{1, 2, 3, 4}, []int{}, []int{}},
	{"all equal", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	{"start equal", []int{1, 2, 5, 5}, []int{1, 2, 6, 6}, []int{1, 2}},
	{"end equal", []int{5, 5, 1, 2}, []int{6, 6, 1, 2}, []int{1, 2}},
	{"middle equal", []int{5, 1, 2, 5}, []int{6, 1, 2, 6}, []int{1, 2}},
	{"multiple equal", []int{1, 2, 5, 5, 5, 6, 7, 8}, []int{1, 2, 4, 4, 6, 7, 8}, []int{1, 2, 6, 7, 8}},
}

func TestLCS(t *testing.T) {
	for _, tc := range lcsTests {
		t.Run(tc.name, func(t *testing.T) {
			got := LCS(tc.inputA, tc.inputB)
			if !slices.Equal(got, tc.want) {
				t.Errorf("LCS(%v, %v) = %v, want %v", tc.inputA, tc.inputB, got, tc.want)
			}
		})
	}
}
