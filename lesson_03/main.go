package main
// Functions in Go


import (
	"math"
	"fmt"
)

func main() {
	// Example of using the calc function
	sum, diff, prod, quot, err := calc(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
		fmt.Println("Difference:", diff)
		fmt.Println("Product:", prod)
		fmt.Println("Quotient:", quot)
	}
	

	// numbers := []int{1, 2, 3, 4, 5}
	// for _, number := range numbers {
	// 	isPrime, err := isPrime(number)
	// 	if err != nil {
	// 		fmt.Printf("number %d : Error: %v\n", number, err)
	// 	} else {
	// 		fmt.Printf("Number %d is prime: %v\n", number, isPrime)
	// 	}
	// }
	// result := add(5, 3)
	// fmt.Println("The sum is:", result)

	// result2, err := divide(10, 2)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("The result of division is:", result2)
	// }

	// result3, result4 := split(17)
	// fmt.Println("Split result:", result3, result4)

}

func add(x int, y int) int {
		return x+y
}


func divide(a, b float64) (float64, error) {
	if b== 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a/b, nil
}


func split(sum int) (x, y int) {
	x = sum*4/9
	y = sum - x
	return
}

func isPrime(n int) (bool, error) {
	if n < 2 {
		return false, fmt.Errorf("number must be greater than >= 2")
	}
	if n == 2 {
		return true, nil
	}
	if n%2 == 0 {
		return false, nil
	}
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func calc(a, b int) (sum, diff, prod, quot int, err error) {
	sum = a + b
	diff = a - b
	prod = a * b
	if b == 0 {
		return 0, 0, 0, 0, fmt.Errorf("division by zero")
	}
	quot = a / b

	return
}
