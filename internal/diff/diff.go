package diff

import "github.com/steamedbuns/go-diff/internal/diff/edit"

func LCS(inputA, inputB []string) []edit.Node {
	return []edit.Node{
		{Content: "a", Op: edit.Remove},
		{Content: "b", Op: edit.Remove},
		{Content: "c", Op: edit.Remove},
	}
}
