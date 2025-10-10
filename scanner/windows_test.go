package scanner

import "testing"

func TestMapPackageName(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Python.Python.3.12", "python"},
		{"Python.Python.3.11", "python"},
		{"Python.Python.3.9", "python"},
		{"Git.Git", "git"},
		{"SomeOther.Package", "SomeOther.Package"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := mapPackageName(tc.input)
			if actual != tc.expected {
				t.Errorf("mapPackageName(%q) = %q; want %q", tc.input, actual, tc.expected)
			}
		})
	}
}