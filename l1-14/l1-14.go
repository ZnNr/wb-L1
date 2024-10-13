package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу,
которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func detectType(variable interface{}) {
	// Получаем тип переменной с помощью reflect
	vType := reflect.TypeOf(variable)

	switch vType.Kind() {
	case reflect.Int:
		fmt.Println("Тип переменной: int")
	case reflect.String:
		fmt.Println("Тип переменной: string")
	case reflect.Bool:
		fmt.Println("Тип переменной: bool")
	case reflect.Chan:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Println("Другой тип переменной:", vType)
	}
}

func main() {
	var intVar int = 42
	var stringVar string = "Hello, World!"
	var boolVar bool = true
	var channelVar chan int = make(chan int)

	// Проверка типов переменных
	detectType(intVar)
	detectType(stringVar)
	detectType(boolVar)
	detectType(channelVar)

	// Пример переменной другого типа
	var floatVar float64 = 3.14
	detectType(floatVar)
}
