package main

import "fmt"

func main() {


	p1 := Person{"Alice", 29}

	p2 := Person{
		Name : "Bob",
		Age	 : 29,
	}

	var p3 Person
	p3.Name = "Charlie"

	fmt.Println(p1, p2, p3)
	p3.Greet()

}

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Я %s, мне %d", p.Name, p.Age)
}

func (p *Person) Birthday() {
	p.Age++
}