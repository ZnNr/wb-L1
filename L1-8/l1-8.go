package main

import (
	"fmt"
	"strconv"
)

/*
Дана переменная int64.
Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var a int64
	var pos int
	var bit int // ввод 0 или 1 корректно работал

	// Ввод натурального числа
	fmt.Print("Введите натуральное число: ")
	_, err := fmt.Scan(&a)
	if err != nil || a < 0 {
		fmt.Println("Необходимо ввести неотрицательное число!", err)
		return
	}

	// Ввод позиции бита
	fmt.Print("Введите позицию бита (начиная с 1): ")
	_, err = fmt.Scan(&pos)
	if err != nil || pos < 1 {
		fmt.Println("Необходимо ввести натуральное число для позиции!", err)
		return
	}

	// Ввод бита (0 или 1)
	fmt.Print("Введите 0 или 1: ")
	_, err = fmt.Scan(&bit)
	if err != nil || (bit != 0 && bit != 1) {
		fmt.Println("Необходимо ввести 0 или 1!", err)
		return
	}

	// Установка bita в соответствующей позиции
	if bit == 1 {
		a |= (1 << (pos - 1)) // Установить i-й бит в 1
	} else {
		a &= ^(1 << (pos - 1)) // Установить i-й бит в 0
	}

	// Вывод результата
	fmt.Printf("Результат: %d (в двоичном: %s)\n", a, strconv.FormatInt(a, 2))
}
