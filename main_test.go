package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := map[string]bool{
		"level":                       true,
		"Level":                       true,
		"A man a plan a canal Panama": true,
		"race a car":                  false,
		"hello":                       false,
		"":                            false,
	}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) 期望 %v 得到 %v", in, want, got)
		}
	}
}
