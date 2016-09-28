/*
Url: [two-sun](https://leetcode.com/problems/two-sum/)
Example:
	```
	Given nums = [2, 7, 11, 15], target = 9,
	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].
	```
Version: 1
*/
package main

import "fmt"

var (
	//nums   = []int{2, 7, 11, 15}
	//target = 9
	nums   = []int{0, 1, 2, 0}
	target = 0
	result = make([]int, 0)
)

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i != j && i < j {
				if nums[i]+nums[j] == target {
					result = append(result, i)
					result = append(result, j)
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	fmt.Println(result)
	return result // result: [0 3]
}

func main() {
	twoSum(nums, target)
}
