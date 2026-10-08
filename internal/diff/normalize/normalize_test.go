package normalize

import (
	"slices"
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

var removePriorityTests = []struct {
	name string
	in   []edit.Node
	want []edit.Node
}{
	{name: "empty", in: []edit.Node{}, want: []edit.Node{}},
	{name: "nil", in: nil, want: []edit.Node{}},
	// {
	// 	name: "only unchanged",
	// 	in:   []edit.Node{edit.Keep("a"), edit.Keep("b")},
	// 	want: []edit.Node{edit.Keep("a"), edit.Keep("b")},
	// },
	// {name: "single add", in: []edit.Node{edit.Add("a")}, want: []edit.Node{edit.Add("a")}},
	// {name: "single remove", in: []edit.Node{edit.Remove("a")}, want: []edit.Node{edit.Remove("a")}},
	// {
	// 	name: "already ordered",
	// 	in:   []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
	// 	want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
	// },
	// {
	// 	name: "edit.Add before remove",
	// 	in:   []edit.Node{edit.Add("x"), edit.Remove("a")},
	// 	want: []edit.Node{edit.Remove("a"), edit.Add("x")},
	// },
	// {
	// 	name: "interleaved edit.Keeps relative order",
	// 	in:   []edit.Node{edit.Add("x"), edit.Remove("a"), edit.Add("y"), edit.Remove("b")},
	// 	want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
	// },
	// {
	// 	name: "block in the middle",
	// 	in:   []edit.Node{edit.Keep("k"), edit.Add("x"), edit.Remove("a"), edit.Keep("m")},
	// 	want: []edit.Node{edit.Keep("k"), edit.Remove("a"), edit.Add("x"), edit.Keep("m")},
	// },
	// {
	// 	name: "blocks at start and end",
	// 	in:   []edit.Node{edit.Add("x"), edit.Remove("a"), edit.Keep("k"), edit.Add("y"), edit.Remove("b")},
	// 	want: []edit.Node{edit.Remove("a"), edit.Add("x"), edit.Keep("k"), edit.Remove("b"), edit.Add("y")},
	// },
	// {
	// 	name: "remove does not cross unchanged line",
	// 	in:   []edit.Node{edit.Add("x"), edit.Keep("k"), edit.Remove("a")},
	// 	want: []edit.Node{edit.Add("x"), edit.Keep("k"), edit.Remove("a")},
	// },
	// {
	// 	name: "only edit.Adds edit.Keep order",
	// 	in:   []edit.Node{edit.Add("x"), edit.Add("y"), edit.Add("z")},
	// 	want: []edit.Node{edit.Add("x"), edit.Add("y"), edit.Add("z")},
	// },
	// {
	// 	name: "only removes edit.Keep order",
	// 	in:   []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Remove("c")},
	// 	want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Remove("c")},
	// },
}

func TestRemovePriority(t *testing.T) {
	for _, tc := range removePriorityTests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemovePriority(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Errorf("RemovePriority(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestRemovePriorityIdempotent(t *testing.T) {
	for _, tc := range removePriorityTests {
		t.Run(tc.name, func(t *testing.T) {
			once := RemovePriority(tc.in)
			twice := RemovePriority(once)
			if !slices.Equal(once, twice) {
				t.Errorf("RemovePriority applied twice to %v = %v, want %v", tc.in, twice, once)
			}
		})
	}
}
