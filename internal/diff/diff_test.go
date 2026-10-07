package diff

import "testing"

func TestPlaceholder(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "Placeholder", want: "hello world"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Placeholder()
			if got != tc.want {
				t.Errorf("Placeholder() = %v, want %v", got, tc.want)
			}
		})
	}
}
