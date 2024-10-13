package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

// Функция, которая переворачивает строку
func reverse(s string) string {
	// Преобразуем строку в срез рун, чтобы корректно обрабатывать символы Unicode
	runes := []rune(s)
	n := len(runes)

	// Переворачиваем РУНЫ
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	return string(runes)
}

func main() {
	// Ввод строки
	input := "главрыба"
	reversed := reverse(input)

	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
