package main

// https://leetcode.com/problems/length-of-last-word

import "strings"

func lengthOfLastWord(s string) int {
	lastword := func(text string) string {
		trimmed := strings.TrimSpace(s)
		splitted := strings.Split(trimmed, " ")

		return splitted[len(splitted)-1]
	}

	return len(lastword(s))
}
