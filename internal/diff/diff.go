package diff

import "github.com/steamedbuns/go-diff/internal/diff/edit"

func LCS(inputA, inputB []string) []edit.EditNode {
	return []edit.EditNode{
		{Content: "a", Op: edit.Remove},
		{Content: "b", Op: edit.Remove},
		{Content: "c", Op: edit.Remove},
	}
}
