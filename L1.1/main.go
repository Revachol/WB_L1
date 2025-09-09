package main

import (
	"fmt"
)

type Human struct {
	Name   string
	Age    int
	Height float64
	Weight float64
}

type Action struct {
	Human
	Activity string
}

func (h *Human) Introduce() {
	fmt.Printf("Привет, я %s, мне %d лет\n", h.Name, h.Age)
}

func (h *Human) CalculateBMI() float64 {
	return h.Weight / (h.Height * h.Height)
}

func (a *Action) Do() {
	fmt.Printf("%s занимается %s\n", a.Name, a.Activity)
}

func (a *Action) ShowInfo() {
	a.Introduce()
	fmt.Printf("ИМТ: %.1f\n", a.CalculateBMI())
	a.Do()
}

func main() {
	person := Human{
		Name:   "Dima",
		Age:    20,
		Height: 1.87,
		Weight: 85,
	}

	action := Action{
		Human:    person,
		Activity: "Programming",
	}

	fmt.Println("=== Методы Human через Action ===")
	action.Introduce()
	fmt.Printf("ИМТ: %.1f\n", action.CalculateBMI())

	fmt.Println("\n=== Методы Action ===")
	action.Do()

	fmt.Println("\n=== Комбинированный метод ===")
	action.ShowInfo()

	fmt.Printf("\nДанные: %s, %d лет, рост: %.2fм, вес: %.1fкг\n",
		action.Name, action.Age, action.Height, action.Weight)
}
