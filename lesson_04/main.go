package main

import "fmt"

func main() {
	p1 := Person{
		Name: "John",
		Age:  30,
	}
	fmt.Println(p1)

	p2 := Person{"Jane", 25}
	fmt.Println(p2)
	
	p3 := Person{}
	p3.Name = "Doe"
	p3.Age = 35
	fmt.Println(p3)

	p4 := Person{"Alice", 0}
	p4.Age = 28
	fmt.Println(p4)

	fmt.Println(p1.Greet())
	p1.Birthday()
	fmt.Println("After birthday:", p1.Age)

}


type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}


func (p *Person) Birthday() {
	p.Age++
}

type Rectangle struct {
	Length int
	Width  int
}
func (r Rectangle) Area() int {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}
func (r *Rectangle) Scale(factor int) {
	r.Length *= factor
	r.Width *= factor
}