package normalize

import (
	"slices"
	"testing"

	"github.com/steamedbuns/go-diff/internal/diff/edit"
)

var editNodeGroupingTests = []struct {
	name string
	in   []edit.Node
	want []edit.Node
}{
	{name: "empty", in: []edit.Node{}, want: []edit.Node{}},
	{name: "nil", in: nil, want: []edit.Node{}},
	{
		name: "only unchanged",
		in:   []edit.Node{edit.Keep("a"), edit.Keep("b")},
		want: []edit.Node{edit.Keep("a"), edit.Keep("b")},
	},
	{name: "single add", in: []edit.Node{edit.Add("a")}, want: []edit.Node{edit.Add("a")}},
	{name: "single remove", in: []edit.Node{edit.Remove("a")}, want: []edit.Node{edit.Remove("a")}},
	{
		name: "already ordered",
		in:   []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
		want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
	},
	{
		name: "add before remove",
		in:   []edit.Node{edit.Add("x"), edit.Remove("a")},
		want: []edit.Node{edit.Remove("a"), edit.Add("x")},
	},
	{
		name: "interleaved keeps relative order",
		in:   []edit.Node{edit.Add("x"), edit.Remove("a"), edit.Add("y"), edit.Remove("b")},
		want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Add("x"), edit.Add("y")},
	},
	{
		name: "block in the middle",
		in:   []edit.Node{edit.Keep("k"), edit.Add("x"), edit.Remove("a"), edit.Keep("m")},
		want: []edit.Node{edit.Keep("k"), edit.Remove("a"), edit.Add("x"), edit.Keep("m")},
	},
	{
		name: "blocks at start and end",
		in:   []edit.Node{edit.Add("x"), edit.Remove("a"), edit.Keep("k"), edit.Add("y"), edit.Remove("b")},
		want: []edit.Node{edit.Remove("a"), edit.Add("x"), edit.Keep("k"), edit.Remove("b"), edit.Add("y")},
	},
	{
		name: "remove does not cross unchanged line",
		in:   []edit.Node{edit.Add("x"), edit.Keep("k"), edit.Remove("a")},
		want: []edit.Node{edit.Add("x"), edit.Keep("k"), edit.Remove("a")},
	},
	{
		name: "only adds keep order",
		in:   []edit.Node{edit.Add("x"), edit.Add("y"), edit.Add("z")},
		want: []edit.Node{edit.Add("x"), edit.Add("y"), edit.Add("z")},
	},
	{
		name: "only removes keep order",
		in:   []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Remove("c")},
		want: []edit.Node{edit.Remove("a"), edit.Remove("b"), edit.Remove("c")},
	},
}

func TestEditNodeGrouping(t *testing.T) {
	for _, tc := range editNodeGroupingTests {
		t.Run(tc.name, func(t *testing.T) {
			got := EditNodeGrouping(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Errorf("ApplyRemoveLinePriority(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestEditNodeGroupingIdempotent(t *testing.T) {
	for _, tc := range editNodeGroupingTests {
		t.Run(tc.name, func(t *testing.T) {
			once := EditNodeGrouping(tc.in)
			twice := EditNodeGrouping(once)
			if !slices.Equal(once, twice) {
				t.Errorf("ApplyRemoveLinePriority applied twice to %v = %v, want %v", tc.in, twice, once)
			}
		})
	}
}
