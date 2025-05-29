package main

import (
	"fmt"
	"os"
	"strconv"
	)  

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as an argument.")
		return
	}
	limitStr := os.Args[1]
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		fmt.Println("Invalid number:", limitStr)
		return
	}

	for i := 0; i < limit; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
	}


// func main() {
// 	// var age int = 25
// 	// name := "Xurshid"
// 	// var isTeacher bool = true

// 	// fmt.Println	(name, "is", age, "years old.")
// 	// fmt.Println("Is", name, "a teacher?", isTeacher)

// 	// var x int = 15
// 	// if x > 10 {
// 	// 	fmt.Println("x is greater than 10")
// 	// } else if x == 10 {
// 	// 	fmt.Println("x is equal to 10")
// 	// } else {
// 	// 	fmt.Println("x is less than 10")
// 	// }


// 	// for i := 0; i<5; i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	// j := 0
// 	// for j < 5 {
// 	// 	fmt.Println(j)
// 	// 	j++
// 	// }

// 	// day := "wed"

// 	// switch day {
// 	// 	case "mon":
// 	// 		fmt.Println("Monday")
// 	// 	case "tue":
// 	// 		fmt.Println("Tuesday")
// 	// 	case "wed":
// 	// 		fmt.Println("Wednesday")
// 	// 	default :
// 	// 		fmt.Println("Unknown day")
// 	// }


// 	var num int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&num)


// 	// if num%3 == 0 && num%5 == 0 {
// 	// 	fmt.Println(num, "is divisible by 3 and 5")
// 	// } else if num%3 == 0 {
// 	// 	fmt.Println(num, "is divisible by 3")
// 	// } else if num%5 == 0 {
// 	// 	fmt.Println(num, "is divisible by 5")
// 	// 	fmt.Println(num, "is over")
// 	// } else {
// 	// 	fmt.Println(num, "is nummber")
// 	// }

// 	// switch {
// 	// case num%3 == 0 && num%5 == 0:
// 	// 	fmt.Println(num, "is divisible by 3 and 5")
// 	// case num%3 == 0:
// 	// 	fmt.Println(num, "is divisible by 3")
// 	// case num%5 == 0:
// 	// 	fmt.Println(num, "is divisible by 5")
// 	// 	fmt.Printf("%d is over\n", num)
// 	// default:
// 	// 	fmt.Println(num, "is nummber")
// 	// }

// 	// for i := 1; i <= 20; i++ {
// 	// 	output := ""
// 	// 	if i%3 ==0 {
// 	// 		output += "Fizz"
// 	// 	}
// 	// 	if i%5 == 0 {
// 	// 		output += "Buzz"
// 	// 	}
// 	// 	if output == "" {
// 	// 		output = fmt.Sprintf("%d", i)
// 	// 	}
// 	// 	fmt.Println(output)
// 	// }
	
// }


