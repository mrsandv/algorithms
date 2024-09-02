package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []TestCase{
		{"", true},
		{"a", true}, {"anita lava la tina", true}, {"Hola mundo", false},
	}
	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			result := IsPalindrome(tc.Input)
			if result != tc.Expected {
				t.Errorf("Expected %t, but got %t for Input: %s", tc.Expected, result, tc.Input)
			}
		})
	}
}
