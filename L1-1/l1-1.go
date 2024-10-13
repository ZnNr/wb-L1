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

// Метод для структуры Human
func (h *Human) Speak() {
	fmt.Printf("Меня зовут %s. Мне %d лет, и я %s.\n", h.Name, h.Age, h.Gender)
}

func (h *Human) CelebrateBirthday() {
	h.Age++
	fmt.Printf("Сегодня мой день рождения! Мне теперь %d лет.\n", h.Age)
}

// Определяем структуру Action, которая включает в себя Human
type Action struct {
	Human
	Activity string
}

// Метод для структуры Action
func (a *Action) Perform() {
	fmt.Printf("%s выполняет действие: %s.\n", a.Name, a.Activity)
}

func main() {
	// Создаем экземпляр Human
	person := Human{Name: "Анна", Age: 30, Gender: "женщина"}
	person.Speak()
	person.CelebrateBirthday()

	// Создаем экземпляр Action
	action := Action{
		Human:    Human{Name: "Иван", Age: 25, Gender: "мужчина"},
		Activity: "бегает",
	}
	action.Speak() // Используем метод из встроенной структуры Human
	action.Perform()
}

/*

Объяснение:
1. Структура Human:
- Содержит поля: Name, Age, Gender.
- Имеет методы Speak и CelebrateBirthday.
2. Структура Action:
- Включает в себя структуру Human, что позволяет ей использовать методы и поля Human.
3. Методы:
- Speak у Action вызывается благодаря встроенному Human.
- Perform — это уникальный метод Action.
Таким образом, вы получаете аналог наследования в Go через композицию и встраивание.
*/
