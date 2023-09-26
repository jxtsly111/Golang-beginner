package main

import "fmt"

func main() {
	fmt.Println(" welcome to functions in golang")
	greeter()

	result := adder(6, 5)

	fmt.Println("Result is : ", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is ", proRes)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(value ...int) int {
	total := 0

	for _, val := range value {
		total = total + val
	}

	return total
}

func greeter() {
	fmt.Println("Akwaaba from golang")

}
