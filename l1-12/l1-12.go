package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	// Исходная последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество
	set := make(map[string]struct{})

	// Заполняем множество с помощью цикла
	for _, str := range strings {
		set[str] = struct{}{} // Добавляем строку в множество
	}

	// Выводим элементы множества
	fmt.Println("Элементы множества:")
	for key := range set {
		fmt.Println(key)
	}
}
