package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:

abcd — true

abCdefAaf — false

aabcd — false
*/

func hasUniqueChars(s string) bool {
	// Создаем множество для хранения уникальных символов
	charSet := make(map[rune]struct{})

	// Приводим строку к нижнему регистру для регистронезависимой проверки
	s = strings.ToLower(s)

	// Перебираем каждый символ в строке
	for _, char := range s {
		// Проверяем, существует ли символ в множестве
		if _, exists := charSet[char]; exists {
			// Если символ уже есть, возвращаем false
			return false
		}
		// Добавляем символ в множество
		charSet[char] = struct{}{}
	}

	// Если ни один символ не повторился, возвращаем true
	return true
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aabcd", "AaCcDd"}

	for _, str := range testStrings {
		result := hasUniqueChars(str)
		fmt.Printf("Строка: %s, Уникальные символы: %t\n", str, result)
	}
}
