package main

func twoSum(numbers []int, target int) []int {
	result := [2]int{}
	if numbers == nil || len(numbers) <= 0 {
		return result[:]
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			result[0] = i + 1
			result[1] = j + 1
			break
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return result[:]
}

func main() {

}
