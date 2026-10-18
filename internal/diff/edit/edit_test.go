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
		input    int
		callFunc func(t int) Node
		want     Node
	}{
		{name: "Keep", input: 0, callFunc: Keep, want: Node{Token: 0, Op: NoChange}},
		{name: "Add", input: 1, callFunc: Add, want: Node{Token: 1, Op: AddLine}},
		{name: "Remove", input: 2, callFunc: Remove, want: Node{Token: 2, Op: RemoveLine}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.callFunc(tc.input)
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
		in:   []Node{Keep(0), Keep(1)},
		want: []Node{Keep(0), Keep(1)},
	},
	{name: "single add", in: []Node{Add(0)}, want: []Node{Add(0)}},
	{name: "single remove", in: []Node{Remove(0)}, want: []Node{Remove(0)}},
	{
		name: "already ordered",
		in:   []Node{Remove(0), Remove(1), Add(2), Add(3)},
		want: []Node{Remove(0), Remove(1), Add(2), Add(3)},
	},
	{
		name: "add before remove",
		in:   []Node{Add(10), Remove(20)},
		want: []Node{Remove(20), Add(10)},
	},
	{
		name: "interleaved keeps relative order",
		in:   []Node{Add(10), Remove(20), Add(30), Remove(40)},
		want: []Node{Remove(20), Remove(40), Add(10), Add(30)},
	},
	{
		name: "block in the middle",
		in:   []Node{Keep(100), Add(10), Remove(20), Keep(200)},
		want: []Node{Keep(100), Remove(20), Add(10), Keep(200)},
	},
	{
		name: "blocks at start and end",
		in:   []Node{Add(10), Remove(20), Keep(100), Add(30), Remove(40)},
		want: []Node{Remove(20), Add(10), Keep(100), Remove(40), Add(30)},
	},
	{
		name: "remove does not cross unchanged line",
		in:   []Node{Add(10), Keep(100), Remove(20)},
		want: []Node{Add(10), Keep(100), Remove(20)},
	},
	{
		name: "only adds keep order",
		in:   []Node{Add(10), Add(20), Add(30)},
		want: []Node{Add(10), Add(20), Add(30)},
	},
	{
		name: "only removes keep order",
		in:   []Node{Remove(10), Remove(20), Remove(30)},
		want: []Node{Remove(10), Remove(20), Remove(30)},
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
