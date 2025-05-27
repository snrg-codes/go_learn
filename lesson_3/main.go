// Function in Go

package main

import "fmt"


func main() {
	result := add(3,5)
	fmt.Println(result)

	result2, err := divide(10, 2)
	if err != nil {
		fmt.Println("Xato", err)
	} else {
		fmt.Println("Natija", result2)
	}

	sum, diff, prod, quot, err := calc(10,0)
	if err != nil {
		fmt.Println("Xato", err)
	} else {
		fmt.Println("Sum:", sum, "Diff:", diff, "Prod:", prod, "Quot:", quot)
	}
}

func add(a int, b int) int {
	return a+b
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("nol ga bo'lib bo'lmaydi")
	}
	return a/b, nil
}

func calc (a,b int) (sum, diff, prod, quot int, err error) {
	sum = a+b
	diff = a-b
	prod = a*b
	if b == 0 {
		return 0,0,0,0, fmt.Errorf("nolga bo'lib bo'lmaydi")
	}
	quot = a/b
	return
}
