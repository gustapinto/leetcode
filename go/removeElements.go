package main

// https://leetcode.com/problems/remove-element

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)

			i -= 1
		}
	}

	return len(nums)
}
