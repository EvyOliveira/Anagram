package main

import "fmt"

func findMajorityElement(nums []int) int {
	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}
	return candidate
}

func main() {
	nums1 := []int{1000, 2, 300, 40, 300, 300, 1000}
	nums2 := []int{2, 1000, 1, 500, 2, 20, 1000}

	majorityElement1 := findMajorityElement(nums1)
	majorityElement2 := findMajorityElement(nums2)

	fmt.Println("majority element in nums1:", majorityElement1)
	fmt.Println("majority element in nums2:", majorityElement2)
}
