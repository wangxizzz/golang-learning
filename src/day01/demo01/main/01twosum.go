package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result [2]int
	if nums == nil || len(nums) == 0 {
		return result[:]
	}
	var m = map[int]int{}
	for i, v := range nums {
		if val, ok := m[target-v]; ok {
			result[0] = i
			result[1] = val
		} else {
			m[v] = i
		}
	}
	return result[:]
}
func main() {
	nums := [...]int{2, 7, 11, 15}
	fmt.Println(twoSum(nums[:], 9))
}
