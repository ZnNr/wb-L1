package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

// binarySearch получает отсортированный срез и искомое значение
func binarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	// Проверяем, попали ли мы на нужный индекс
	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1 // Если элемент не найден, возвращаем -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d.\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве.\n", target)
	}
}
