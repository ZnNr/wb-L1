package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

//алгоритм быстрой сортировки (quicksort) с использованием встроенных методов.

// quickSort выполняет быструю сортировку массива с помощью рекурсии
func quickSort(array []int) []int {
	if len(array) < 2 {
		return array // массив из 0 или 1 элемента уже отсортирован
	}

	pivot := array[len(array)/2] // выбираем опорный элемент
	left := []int{}
	right := []int{}
	var equal []int // массив для элементов, равных опорному

	for _, value := range array {
		if value < pivot {
			left = append(left, value) // элементы меньше опорного
		} else if value > pivot {
			right = append(right, value) // элементы больше опорного
		} else {
			equal = append(equal, value) // элементы равны опорному
		}
	}

	// рекурсивное применение быстрой сортировки и объединение массивов
	return append(append(quickSort(left), equal...), quickSort(right)...)
}

func main() {
	// Пример массива для сортировки
	array := []int{34, 7, 23, 32, 5, 62}
	fmt.Println("Исходный массив:", array)

	// Сортировка массива
	sortedArray := quickSort(array)
	fmt.Println("Отсортированный массив:", sortedArray)
}
