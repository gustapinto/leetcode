package main

// https://leetcode.com/problems/jewels-and-stones

func numJewelsInStones(jewels string, stones string) int {
	counter := 0

	contains := func(stone rune) bool {
		for _, jewel := range jewels {
			if jewel != stone {
				continue
			}

			return true
		}

		return false
	}

	for _, stone := range stones {
		if contains(stone) {
			counter += 1
		}
	}

	return counter
}
