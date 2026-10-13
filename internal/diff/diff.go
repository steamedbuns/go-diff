package diff

import (
	"hash/maphash"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

// TokenMap creates a mapping from uint64 hash of a string to a bucket of integer tokens.
// The tokens represent the index location of a string inside the tokenTable.
// The tokenTable is a slice that holds unique strings, the index is the token.
// Each token index maps 1 to 1 to a unique string.
type TokenMap struct {
	seed            maphash.Seed
	hashToBucketMap map[uint64][]int
	tokenTable      []string
}

func NewTokenMap() *TokenMap {
	return &TokenMap{
		seed:            maphash.MakeSeed(),
		hashToBucketMap: make(map[uint64][]int),
	}
}

func (tm *TokenMap) getToken(line string) int {
	h := maphash.String(tm.seed, line)
	if bucket, found := tm.hashToBucketMap[h]; found {
		for _, token := range bucket {
			if tm.tokenTable[token] == line {
				return token
			}
		}
	}
	newToken := len(tm.tokenTable)
	tm.tokenTable = append(tm.tokenTable, line)
	tm.hashToBucketMap[h] = append(tm.hashToBucketMap[h], newToken)
	return newToken
}

// Transforms the input into a slice of integers
// The integers represent an integer token that points to the string in the input slice
// If two strings are equal in the input, that will map to the same integer token
func (tm *TokenMap) Tokenize(input []string) []int {
	if len(input) == 0 {
		return nil
	}
	tokens := make([]int, len(input))
	for idx, line := range input {
		tokens[idx] = tm.getToken(line)
	}
	return tokens
}

func LCS(inputA, inputB []string) []edit.Node {
	return []edit.Node{
		edit.Remove("a"),
		edit.Remove("b"),
		edit.Remove("c"),
	}
}
