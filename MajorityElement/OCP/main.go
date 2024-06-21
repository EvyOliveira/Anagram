package main

import (
	"errors"
	"fmt"
)

type MajorityChecker interface {
	FindMajorityElement(num []int64) (int64, error)
}

type MajorityElement struct{}

func (h *MajorityElement) FindMajorityElement(nums []int64) (int64, error) {
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
	return candidate, nil
}

func findMajorityElement(checker MajorityChecker, nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	return checker.FindMajorityElement(nums)
}

func isAnEmptySlice(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	return 1, nil
}

func main() {
	nums1 := []int64{2, 40, 300, 40, 1000, 300}
	nums2 := []int64{1000, 2, 300, 300, 300, 40, 1000, 1000}

	checker := &MajorityElement{}
	_, firstSliceError := isAnEmptySlice(nums1)
	if firstSliceError != nil {
		fmt.Println("error checking nums1:", firstSliceError)
	}

	_, secoundSliceError := isAnEmptySlice(nums2)
	if secoundSliceError != nil {
		fmt.Println("error checking nums2:", secoundSliceError)
	}

	majorityElement1, err := findMajorityElement(checker, nums1)
	if err != nil {
		fmt.Println("error finding majority element in nums1:", err)
	} else {
		fmt.Println("majority element in nums1:", majorityElement1)
	}

	majorityElement2, err := findMajorityElement(checker, nums2)
	if err != nil {
		fmt.Println("error finding majority element in nums2:", err)
	} else {
		fmt.Println("majority element in nums2:", majorityElement2)
	}
}
