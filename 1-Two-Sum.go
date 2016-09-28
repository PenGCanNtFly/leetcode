/*
* version := 1
 */
package main

var (
	//nums   = []int{2, 7, 11, 15}
	//target = 9
	nums   = []int{0, 1, 2, 0}
	target = 0
	result = make([]int, 0)
)

func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j < len(nums); j++ {
			if i != j && i < j {
				if nums[i]+nums[j] == target {
					result = append(result, i, j)
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	return result
}

func main() {
	twoSum(nums, target)
}
