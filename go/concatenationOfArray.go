package main

func getConcatenation(nums []int) []int {
    newNums := nums
        
    for _, i := range nums {
        newNums = append(newNums, i)
    }
    
    return newNums
}
