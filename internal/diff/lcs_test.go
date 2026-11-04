package diff_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff"
	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

// render turns a script into readable lines, prefixed " ", "+" or "-",
// so tests compare text rather than token ids.
func render(s diff.Script) []string {
	out := make([]string, len(s.Nodes))
	for i, n := range s.Nodes {
		var prefix string
		switch n.Op {
		case edit.NoChange:
			prefix = " "
		case edit.AddLine:
			prefix = "+"
		case edit.RemoveLine:
			prefix = "-"
		default:
			prefix = "?"
		}
		out[i] = prefix + s.Text(n)
	}
	return out
}

// reconstruct rebuilds both inputs from a script: unchanged and removed
// lines give the old file, unchanged and added lines give the new file.
func reconstruct(s diff.Script) (oldLines, newLines []string) {
	for _, n := range s.Nodes {
		switch n.Op {
		case edit.NoChange:
			oldLines = append(oldLines, s.Text(n))
			newLines = append(newLines, s.Text(n))
		case edit.RemoveLine:
			oldLines = append(oldLines, s.Text(n))
		case edit.AddLine:
			newLines = append(newLines, s.Text(n))
		}
	}
	return oldLines, newLines
}

var lcsScriptTests = []struct {
	name     string
	oldLines []string
	newLines []string
	want     []string
}{
	{"both empty", []string{}, []string{}, []string{}},
	{"different single line", []string{"foo"}, []string{"bar"}, []string{"-foo", "+bar"}},
	{"only new lines", []string{}, []string{"a", "b"}, []string{"+a", "+b"}},
	{"only old lines", []string{"a", "b"}, []string{}, []string{"-a", "-b"}},
	{"shared lines", []string{"a", "b"}, []string{"b", "c"}, []string{"-a", " b", "+c"}},
	{"duplicate lines", []string{"a", "a"}, []string{"a"}, []string{"-a", " a"}},
}

func TestLCS(t *testing.T) {
	for _, tc := range lcsScriptTests {
		t.Run(tc.name, func(t *testing.T) {
			got := render(diff.LCS(tc.oldLines, tc.newLines))
			if !slices.Equal(got, tc.want) {
				t.Errorf("LCS(%q, %q) =\ngot  %q\nwant %q", tc.oldLines, tc.newLines, got, tc.want)
			}
		})
	}
}

func FuzzLCS(f *testing.F) {
	f.Add("", "")
	f.Add("foo", "bar")
	f.Add("a\nb\nc", "a\nc\nd")
	f.Add("a\na\na", "a")
	f.Fuzz(func(t *testing.T, a, b string) {
		oldLines, newLines := strings.Split(a, "\n"), strings.Split(b, "\n")
		gotOld, gotNew := reconstruct(diff.LCS(oldLines, newLines))
		if !slices.Equal(gotOld, oldLines) {
			t.Errorf("script rebuilds old as %q, want %q", gotOld, oldLines)
		}
		if !slices.Equal(gotNew, newLines) {
			t.Errorf("script rebuilds new as %q, want %q", gotNew, newLines)
		}
	})
}
