package main

import (
	"errors"
	"fmt"
)

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

func isAnEmptySlice(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}
	return 1, nil
}

func main() {
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

	majorityElement1 := findMajorityElement(nums1)
	majorityElement2 := findMajorityElement(nums2)

	fmt.Println("majority element in nums1:", majorityElement1)
	fmt.Println("majority element in nums2:", majorityElement2)
}
