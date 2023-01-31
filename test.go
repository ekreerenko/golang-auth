package main

import (
	"fmt"
	"golang-auth/computation"
)

func main() {

	intCh := make(chan int)

	var value int

	fmt.Scanln(&value)
	go computation.FactorialGorutene(value, intCh)

	fmt.Println(<-intCh)

	fmt.Println("The End")

	// intCh := make(chan int)

	// go func() {
	// 	fmt.Println("Go routine starts")
	// 	intCh <- 5 // блокировка, пока данные не будут получены функцией main
	// }()
	// fmt.Println("extract", <-intCh) // получение данных из канала
	// fmt.Println("The End2")
}
