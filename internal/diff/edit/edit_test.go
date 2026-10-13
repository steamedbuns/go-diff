package edit

import (
	"slices"
	"testing"
)

func TestLineOperationString(t *testing.T) {
	tests := []struct {
		name string
		in   Op
		want string
	}{
		{name: "NoChange", in: NoChange, want: "NoChange"},
		{name: "Add", in: AddLine, want: "AddLine"},
		{name: "Remove", in: RemoveLine, want: "RemoveLine"},
		{name: "Default", in: Op(12), want: "Unknown LineOperation(12)"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.in.String()
			if got != tc.want {
				t.Errorf("LineOperation.String(%v) = %v, want %v", int(tc.in), got, tc.want)
			}
		})
	}
}

func TestNodeHelperFunction(t *testing.T) {
	tests := []struct {
		name     string
		callFunc func(s string) Node
		want     Node
	}{
		{name: "Keep", callFunc: Keep, want: Node{Content: "Keep", Op: NoChange}},
		{name: "Add", callFunc: Add, want: Node{Content: "Add", Op: AddLine}},
		{name: "Remove", callFunc: Remove, want: Node{Content: "Remove", Op: RemoveLine}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.callFunc(tc.name)
			if got != tc.want {
				t.Errorf("%v(%v) = %v, want %v", tc.name, tc.name, got, tc.want)
			}
		})
	}
}

var normalizeTests = []struct {
	name string
	in   []Node
	want []Node
}{
	{name: "empty", in: []Node{}, want: []Node{}},
	{name: "nil", in: nil, want: []Node{}},
	{
		name: "only unchanged",
		in:   []Node{Keep("a"), Keep("b")},
		want: []Node{Keep("a"), Keep("b")},
	},
	{name: "single add", in: []Node{Add("a")}, want: []Node{Add("a")}},
	{name: "single remove", in: []Node{Remove("a")}, want: []Node{Remove("a")}},
	{
		name: "already ordered",
		in:   []Node{Remove("a"), Remove("b"), Add("x"), Add("y")},
		want: []Node{Remove("a"), Remove("b"), Add("x"), Add("y")},
	},
	{
		name: "add before remove",
		in:   []Node{Add("x"), Remove("a")},
		want: []Node{Remove("a"), Add("x")},
	},
	{
		name: "interleaved keeps relative order",
		in:   []Node{Add("x"), Remove("a"), Add("y"), Remove("b")},
		want: []Node{Remove("a"), Remove("b"), Add("x"), Add("y")},
	},
	{
		name: "block in the middle",
		in:   []Node{Keep("k"), Add("x"), Remove("a"), Keep("m")},
		want: []Node{Keep("k"), Remove("a"), Add("x"), Keep("m")},
	},
	{
		name: "blocks at start and end",
		in:   []Node{Add("x"), Remove("a"), Keep("k"), Add("y"), Remove("b")},
		want: []Node{Remove("a"), Add("x"), Keep("k"), Remove("b"), Add("y")},
	},
	{
		name: "remove does not cross unchanged line",
		in:   []Node{Add("x"), Keep("k"), Remove("a")},
		want: []Node{Add("x"), Keep("k"), Remove("a")},
	},
	{
		name: "only adds keep order",
		in:   []Node{Add("x"), Add("y"), Add("z")},
		want: []Node{Add("x"), Add("y"), Add("z")},
	},
	{
		name: "only removes keep order",
		in:   []Node{Remove("a"), Remove("b"), Remove("c")},
		want: []Node{Remove("a"), Remove("b"), Remove("c")},
	},
}

func TestNormalize(t *testing.T) {
	for _, tc := range normalizeTests {
		t.Run(tc.name, func(t *testing.T) {
			got := Normalize(tc.in)
			if !slices.Equal(got, tc.want) {
				t.Errorf("Normalize(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func TestNormalizeIdempotent(t *testing.T) {
	for _, tc := range normalizeTests {
		t.Run(tc.name, func(t *testing.T) {
			once := Normalize(tc.in)
			twice := Normalize(once)
			if !slices.Equal(once, twice) {
				t.Errorf("Normalize applied twice to %v = %v, want %v", tc.in, twice, once)
			}
		})
	}
}
