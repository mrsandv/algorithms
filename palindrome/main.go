package main

import (
	"fmt"
	"strings"
	"testing"
)

func main() {
	arr := []TestCase{{Input: "ana", Expected: true}, {Input: "PeS%rro", Expected: false}, {Input: "anita lava la tina", Expected: true}}
	for i, tc := range arr {
		result := IsPalindrome(tc.Input)
		if result != tc.Expected {
			fmt.Printf("Case: %d, failed: Input: %s, Expected: %t, Got %t\n", i+1, tc.Input, tc.Expected, result)
		} else {
			fmt.Printf("Case: %d, passed: Input: %s\n", i+1, tc.Input)
		}
	}
}

type TestCase struct {
	Input    string
	Expected bool
}

func IsPalindrome(s string) bool {
	// First let's go to normalize the sentence
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	// let's define the extremes of the sentence to run from each side and compare the chars
	left, right := 0, len(s)-1
	for left < right {
		// For an advanced version check the alfanumerics chars, and ignore it
		// Ignore non alfanumeric chars from left
		// for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
		// 	left++
		// }
		// Ignore non alfanumeric chars from right
		// for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
		// 	left--
		// }
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

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
