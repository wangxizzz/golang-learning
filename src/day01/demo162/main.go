package main

import "fmt"

func findPeakElementleetocde162(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	count := len(nums)
	for i, _ := range nums {
		if (i == count-1) && nums[i] > nums[i-1] {
			return i
		}
		if i > 0 && nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}
func main() {
	var nums = [...]int{1, 2, 3, 1}
	fmt.Println(findPeakElementleetocde162(nums[:]))
}
