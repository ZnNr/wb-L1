package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

/*
Чтобы поменять местами два числа без использования временной переменной,
можно использовать два основных подхода: через арифметические операции или
с помощью побитового исключающего ИЛИ (XOR).
*/

// Побитовое исключающее ИЛИ (XOR)
func main() {
	a := 5
	b := 10

	fmt.Println("До обмена:", a, b)

	// Обмен значениями без временной переменной с использованием XOR
	a = a ^ b // a теперь 15 (5 XOR 10)
	b = a ^ b // b теперь 5 (15 XOR 10)
	a = a ^ b // a теперь 10 (15 XOR 5)

	fmt.Println("После обмена:", a, b)
}

//Арифметические операции

//func main() {
//	a := 5
//	b := 10
//
//	fmt.Println("До обмена:", a, b)
//
//	// Обмен значениями без временной переменной
//	a = a + b // a теперь 15 (5 + 10)
//	b = a - b // b теперь 5 (15 - 10)
//	a = a - b // a теперь 10 (15 - 5)
//
//	fmt.Println("После обмена:", a, b)
//}
