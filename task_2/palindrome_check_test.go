package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Simple Single Word",
			input:    "racecar",
			expected: true,
		},
		{
			name:     "Phrase with Punctuation and Spaces",
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Mixed Case Word",
			input:    "Madam",
			expected: true,
		},
		{
			name:     "Non-Palindrome String",
			input:    "hello world",
			expected: false,
		},
		{
			name:     "Numerical Palindrome",
			input:    "12321",
			expected: true,
		},
		{
			name:     "Empty or Blank String",
			input:    "   ",
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("isPalindrome(%q) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
