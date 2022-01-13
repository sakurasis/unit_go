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

/**
twoSum  return indices of the two numbers such that they add up to a specific target.
		steps:
		1.for each element, we always find the another number from which get surplus number, and returns
		the index when found it.
		2. when it was not found, it wound store to a new map util to scan another one the loop.

[]int   the indices is satisfied to the condition which the sum of two numbers to equal the target number.
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		next := target - nums[i]
		// if the new map contain the key, go back to the array.
		if _, y := m[next]; y {
			return []int{m[next], i}
		}
		// return the index to the new map.
		m[nums[i]] = i
	}
	fmt.Printf("not found it.....")
	return nil
}
