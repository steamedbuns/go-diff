package diff

import (
	"hash/maphash"
	"slices"
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

func LCS(inputA, inputB []int) []int {
	lcsGrid := make([][]int, len(inputA)+1)
	for i := range lcsGrid {
		lcsGrid[i] = make([]int, len(inputB)+1)
	}

	for i := 1; i <= len(inputA); i++ {
		for j := 1; j <= len(inputB); j++ {
			if inputA[i-1] == inputB[j-1] {
				lcsGrid[i][j] = lcsGrid[i-1][j-1] + 1
			} else {
				lcsGrid[i][j] = max(lcsGrid[i][j-1], lcsGrid[i-1][j])
			}
		}
	}

	lcs := []int{}
	i, j := len(inputA), len(inputB)
	for i > 0 && j > 0 {
		if inputA[i-1] == inputB[j-1] {
			lcs = append(lcs, inputA[i-1])
			i--
			j--
		} else if lcsGrid[i-1][j] > lcsGrid[i][j-1] {
			i--
		} else {
			j--
		}
	}
	slices.Reverse(lcs)
	return lcs
}
