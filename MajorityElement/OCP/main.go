package main

import (
	"errors"
	"fmt"
)

type MajorityChecker interface {
	CheckMajority(arr []int64, candidate int64) bool
}

func findMajorityElement(nums []int64) int64 {
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

type SimpleMajorityChecker struct{}

func (s *SimpleMajorityChecker) CheckMajority(arr []int64, candidate int64) bool {
	count := int64(0)

	for _, num := range arr {
		if num == candidate {
			count++
			if count > int64(len(arr)/2) {
				return true
			}
		}
	}
	return count > s.countElement(arr, candidate)/2
}

func (s *SimpleMajorityChecker) countElement(arr []int64, element int64) int64 {
	count := int64(0)

	for _, num := range arr {
		if num == element {
			count++
		}
	}
	return count
}

func isAnEmptySlice(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	return 1, nil
}

func main() {
	checker := &SimpleMajorityChecker{}

	nums1 := []int64{1000, 2, 300, 40, 300, 300, 1000}
	nums2 := []int64{1000, 2, 300, 300, 40, 1000}

	_, firstSliceError := isAnEmptySlice(nums1)
	if firstSliceError != nil {
		fmt.Println("error checking nums1:", firstSliceError)
	}

	_, secoundSliceError := isAnEmptySlice(nums2)
	if secoundSliceError != nil {
		fmt.Println("error checking nums2:", secoundSliceError)
	}

	if checker.CheckMajority(nums1, nums1[0]) {
		fmt.Println("majority element in nums1:", nums1[0])
	} else {
		fmt.Println("no majority element in nums1")
	}

	if checker.CheckMajority(nums2, nums2[0]) {
		fmt.Println("majority element in nums2:", nums2[0])
	} else {
		fmt.Println("no majority element in nums2")
	}
}
