package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	weight  float64
	IsAdult bool
}

func New(name string, age int, weight float64) *Person {
	isAdult := age > 18
	return &Person{
		Name:    name,
		Age:     age,
		weight:  weight,
		IsAdult: isAdult,
	}
}

func main() {
	p := Person{Name: "Sudipto", Age: 26, weight: 72, IsAdult: true}
	p2 := New("sudipto", 26, 72)
	fmt.Println(p)
	fmt.Println(p2)
}
func (p *Person) SetName(name string) {

}
