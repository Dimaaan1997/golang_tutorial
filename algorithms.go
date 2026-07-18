package main

import (
	"fmt"
)

// Задача на слияние двух массиовов
func mergeNew(nums1, nums2 []int) []int {
	result := make([]int, 0, len(nums1)+len(nums2))

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}

	result = append(result, nums1[i:]...)
	result = append(result, nums2[j:]...)

	return result
}

// Задача на удаление дубликатов
func removeDuplicates(nums []int) []int {

	if len(nums) == 0 {
		return nil
	}

	result := []int{}

	i, j := 0, 1

	result = append(result, nums[i])

	for j < len(nums) {
		if nums[i] != nums[j] {
			result = append(result, nums[j])
			i = j
		}

		j++

	}
	return result
}

// Задача на удаление дубликатов без создания нового списка
func removeDuplicatesInPlace(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)

	return i + 1

}

// Задача на сдвиг нулей право
func moveZeroes(nums []int) {
	i := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
