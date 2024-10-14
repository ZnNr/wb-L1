package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

// Определяем структуру Human
type Human struct {
	Name   string
	Age    int
	Gender string
}

// Метод для структуры Human, метод представляет чела
func (h *Human) Introduce() {
	fmt.Printf("Меня зовут %s. Мне %d лет, и я %s.\n", h.Name, h.Age, h.Gender)
}

func (h *Human) CelebrateBirthday() {
	h.Age++
	fmt.Printf("Сегодня мой день рождения! Теперь мне %d лет.\n", h.Age)
}

// Определяем структуру Action, которая включает (встраивает) в себя Human
type Action struct {
	Human
	Activity string
}

// Метод для структуры Action
func (a *Action) Perform() {
	fmt.Printf("%s выполняет действие: %s.\n", a.Name, a.Activity)
}

func main() {
	// экземпляр Human
	person := Human{Name: "Полина", Age: 30, Gender: "женщина"}
	person.Introduce()
	person.CelebrateBirthday()

	// экземпляр Action
	action := Action{
		Human:    Human{Name: "Джонни", Age: 25, Gender: "мужчина"},
		Activity: "пресс качать",
	}
	action.Introduce() // Вызываем метод Introduce структуры Action (унаследованный из Human)
	action.Perform()   // Вызываем метод Execute структуры Action
}
