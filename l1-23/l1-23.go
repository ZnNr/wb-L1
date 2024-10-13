package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

// removeElement удаляет элемент по индексу i из слайса
func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		fmt.Println("Индекс вне диапазона.")
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	// Исходный слайс
	slice := []int{1, 2, 3, 4, 5}

	// Индекс для удаления
	indexToRemove := 2 // Удаляем элемент со значением 3

	// Вывод исходного слайса
	fmt.Println("Первоначальный слайс:", slice)

	// Удаление элемента
	newSlice := removeElement(slice, indexToRemove)

	// Вывод измененного слайса
	fmt.Println("Слайс после удаления элемента:", newSlice)
}
