package computation

import "fmt"

func Factorial(n int) int {

	var result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func FactorialGorutene(number int, ch chan<- int) {
	if number < 1 {
		fmt.Println("Invalid input number")
		return
	}

	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}

	fmt.Println(number, " - ", result)

	ch <- result
}
