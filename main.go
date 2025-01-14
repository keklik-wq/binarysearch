package main

import (
	"fmt"
)

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 7

	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден.\n", target)
	}
}
