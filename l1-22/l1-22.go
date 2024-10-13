package main

import (
	"fmt"
	"log"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a и b,
значение которых > 2^20.
*/

func main() {
	// Задать переменные a и b
	a := 2200000 // Значение a > 2^20
	b := 2400000 // Значение b > 2^20

	// Проверка условий
	if a <= (1<<20) || b <= (1<<20) {
		log.Fatal("Значения a и b должны быть больше 2^20.")
	}

	// Сложение
	sum := a + b
	fmt.Printf("Сумма a и b: %d + %d = %d\n", a, b, sum)

	// Вычитание
	diff := a - b
	fmt.Printf("Разность a и b: %d - %d = %d\n", a, b, diff)

	// Умножение
	product := a * b
	fmt.Printf("Произведение a и b: %d * %d = %d\n", a, b, product)

	// Деление
	if b != 0 {
		quotient := float64(a) / float64(b)
		fmt.Printf("Частное a и b: %d / %d = %.2f\n", a, b, quotient)
	} else {
		fmt.Println("Ошибка: деление на ноль.")
	}
}
