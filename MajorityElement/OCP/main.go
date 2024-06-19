package main

import "fmt"

type MajorityChecker interface {
	CheckMajority(arr []int, candidate int) bool
}

func findMajorityElement(nums []int, checker MajorityChecker) (int, bool) {
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

	if checker.CheckMajority(nums, candidate) {
		return candidate, true
	}
	return 0, false
}

type SimpleMajorityChecker struct{}

func (s *SimpleMajorityChecker) CheckMajority(arr []int, candidate int) bool {
	count := 0

	for _, num := range arr {
		if num == candidate {
			count++
			if count > len(arr)/2 {
				return true
			}
		}
	}
	return count > s.countElement(arr, candidate)/2
}

func (s *SimpleMajorityChecker) countElement(arr []int, element int) int {
	count := 0

	for _, num := range arr {
		if num == element {
			count++
		}
	}
	return count
}

func main() {
	checker := &SimpleMajorityChecker{}

	nums1 := []int{1000, 2, 300, 40, 300, 300, 1000}
	nums2 := []int{2, 1000, 1, 500, 2, 20, 1000}

	majorityElement1, isMajority1 := findMajorityElement(nums1, checker)
	majorityElement2, isMajority2 := findMajorityElement(nums2, checker)

	if isMajority1 {
		fmt.Println("majority element in nums1:", majorityElement1)
	} else {
		fmt.Println("no majority element in nums1")
	}

	if isMajority2 {
		fmt.Println("majority element in nums2:", majorityElement2)
	} else {
		fmt.Println("no majority element in nums2")
	}
}
