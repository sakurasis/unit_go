package main

import "fmt"

func main() {
	nums := []int{1, 7, 7, 2}
	target := 9
	sum := twoSum(nums, target)
	for i := 0; i < len(sum); i++ {
		fmt.Printf("nums[%d]", sum[i])
		if i != len(sum)-1 {
			fmt.Printf(" + ")
		} else {
			fmt.Printf(" = %d + %d = %d", nums[sum[0]], nums[sum[i]], target)
		}
	}

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		next := target - nums[i]
		if _, y := m[next]; y {
			return []int{m[next], i}
		}
		m[nums[i]] = i
	}
	fmt.Printf("not found it.....")
	return nil
}
