package main

import (
	"fmt"
)

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify) {
	fmt.Println(s.ToString())
}

func main() {
	vs := []Stringify{
		&Person{Name: "Sam", Age: 20},
		&Car{Number: "001-0001", Model: "A001"},
	}
	for _, v := range vs {
		fmt.Println(v.ToString())
	}

	Println(&Person{Name: "John", Age: 20})
	Println(&Car{Number: "002-0002", Model: "B002"})
}
