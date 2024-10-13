package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.

Пример: «snow dog sun — sun dog snow».
*/

// Функция, которая переворачивает слова в строке
func reverseWords(s string) string {
	// Разбиваем строку на слова
	words := strings.Fields(s)
	n := len(words)

	// Переворачиваем слова
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}

	// Собираем перевернутые слова обратно в строку
	return strings.Join(words, " ")
}

func main() {
	// Ввод строки
	input := "snow dog sun"
	reversed := reverseWords(input)

	fmt.Printf("Исходная строка: \"%s\"\n", input)
	fmt.Printf("Перевернутая строка: \"%s\"\n", reversed)
}
