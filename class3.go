package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	weight float64
}

func New(name string, age int, weight float64) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		weight: weight,
	}
}

func main() {
	p := New("sudipto", 26, 72)
	fmt.Println(*p)
}
