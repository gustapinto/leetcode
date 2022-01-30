package main

// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	reverse := func(n int) int {
		reversed := 0
		for n > 0 {
			r := n % 10
			reversed *= 10
			reversed += r
			n /= 10
		}
		return reversed
	}

	return reverse(x) == x
}
