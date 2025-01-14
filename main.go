package main

import (
	"fmt"
	"sync"
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
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 25, 90, 110, 250, 252, 254, 256, 258}
	targets := []int{7, 11, 20, 1, 258}

	results := make(chan string)

	var wg sync.WaitGroup

	for _, target := range targets {
		wg.Add(1)
		go func(target int) {
			defer wg.Done()
			result := BinarySearch(arr, target)
			if result != -1 {
				results <- fmt.Sprintf("Элемент %d найден на индексе %d.", target, result)
			} else {
				results <- fmt.Sprintf("Элемент %d не найден.", target)
			}
		}(target)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
