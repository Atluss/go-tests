package other

import "testing"

func TestCheckBrackets(t *testing.T) {
	// define test cases.
	tests := []struct {
		sentence string
		expected bool
	}{
		{"{[]()}", true},
		{"{[(])}", false},
		{"{[}", false},
		{"}}{{[}}", false},
	}

	for _, tt := range tests {
		result := CheckBrackets(tt.sentence)
		if tt.expected != result {
			t.Error("fail")
		}
	}
}
