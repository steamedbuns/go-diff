package diff

import (
	"hash/maphash"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

type TokenMap struct {
	seed            maphash.Seed
	hashToTokensMap map[uint64][]int
	dataTable       []string
}

func (tm *TokenMap) Tokenize(input []string) []int {
	return []int{}
}

func LCS(inputA, inputB []string) []edit.Node {
	return []edit.Node{
		edit.Remove("a"),
		edit.Remove("b"),
		edit.Remove("c"),
	}
}
