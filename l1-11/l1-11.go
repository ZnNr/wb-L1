package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

// Функция для нахождения пересечения двух множеств
func intersect(set1, set2 map[int]struct{}) map[int]struct{} {
	intersection := make(map[int]struct{})
	for key := range set1 {
		if _, exists := set2[key]; exists {
			intersection[key] = struct{}{}
		}
	}
	return intersection
}

func main() {
	// Создаем два множества
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	// Заполняем первое множество с помощью цикла
	for i := 1; i <= 10; i++ {
		set1[i] = struct{}{}
	}

	// Заполняем второе множество с помощью цикла
	for i := 5; i <= 15; i++ {
		set2[i] = struct{}{}
	}

	// Нахождение пересечения
	result := intersect(set1, set2)

	// Вывод результата
	fmt.Println("Пересечение множеств:")
	for key := range result {
		fmt.Println(key)
	}
}
