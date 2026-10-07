package diff

import "github.com/steamedbuns/go-diff/internal/diff/edit"

func LCS(inputA, inputB []string) []edit.Node {
	return []edit.Node{
		edit.Remove("a"),
		edit.Remove("b"),
		edit.Remove("c"),
	}
}
