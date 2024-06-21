package main

import (
	"errors"
	"fmt"
	"sort"
)

type MajorityChecker interface {
	FindMajorityElement(num []int64) (int64, error)
}

type HashTableMajorityElement struct{}

func (h *HashTableMajorityElement) FindMajorityElement(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

	elements := make(map[int64]int64)
	for _, num := range nums {
		elements[num]++
		if elements[num] > int64(len(nums)/2) {
			return num, nil
		}
	}
	return -1, errors.New("majority element not found")
}

type SortingMajorityElement struct{}

func (s *SortingMajorityElement) FindMajorityElement(nums []int64) (int64, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

	sortedNums := make([]int64, len(nums))
	copy(sortedNums, nums)
	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i] < sortedNums[j]
	})

	mid := int64(len(sortedNums) / 2)
	count := 1

	for i := mid + 1; i < int64(len(sortedNums)); i++ {
		if sortedNums[i] == sortedNums[mid] {
			count++
		}
	}
	if count > len(nums)/2 {
		return sortedNums[mid], nil
	}

	return sortedNums[mid], nil
}

func isAnEmptySlice(finder MajorityChecker, nums []int64) (int64, error) {
	return finder.FindMajorityElement(nums)
}

func main() {
	nums1 := []int64{2, 40, 300, 40, 1000, 300}
	nums2 := []int64{1000, 2, 300, 40, 1000, 1000}

	finder := &SortingMajorityElement{}

	result1, err := isAnEmptySlice(finder, nums1)
	if err != nil {
		fmt.Println("error for nums1:", err)
	} else {
		fmt.Println("majority element in:", nums1, ":", result1)
	}

	result2, err := isAnEmptySlice(finder, nums2)
	if err != nil {
		fmt.Println("error for nums2:", err)
	} else {
		fmt.Println("majority element in:", nums2, ":", result2)
	}
}
