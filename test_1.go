package main

import (
	"fmt"
)

func add(numbers ...int) (s int) {
	var sum = 0

	for _, number := range numbers {
		sum += number
	}
	s = sum

	fmt.Println("sum = ", sum)
	return
}

type Stream interface {
	read() string
	write(string)
	close()
}

func writeStream(stream Stream, text string) {
	stream.write(text)
}

func closeStream(stream Stream) {
	stream.close()
}

type File struct {
	text string
}

type Folder struct{}

func (f *File) read() string {
	return f.text
}

func (f *File) write(msg string) {
	f.text = msg
	fmt.Println("Запись в файл строки", msg)
}

func (f *File) close() {
	fmt.Println("Файл закрыт")
}

func (f *Folder) close() {
	fmt.Println("Папка закрыта")
}

func main() {
	myFile := &File{}
	myFolder := &Folder{}

	writeStream(myFile, "test message")

	closeStream(myFile)

	myFolder.close()

	// fmt.Println(computation.Factorial(9))

	// var tom = person{name: "Tom", age: 24}

	// var tomPointer *person = &tom

	// fmt.Println("before", tom.age)
	// tomPointer.updateAge(33)
	// fmt.Println("after", tom.age)

	// tom.print()
	// tom.eat("борщ с капустой, но не красный")

	// var x int = 4       // определяем переменную
	// var p *int          // определяем указатель
	// p = &x              // указатель получает адрес переменной
	// fmt.Println(p, x)      // значение самого указателя - адрес переменной x

	// p := new(int)
	// fmt.Println("Value:", *p, p) // Value: 0 - значение по умолчанию
	// *p = 8                       // изменяем значение
	// fmt.Println("Value:", *p, p) // Value: 8
	// var hello string = "some test string"
	// fmt.Println(hello)

	// const (
	//   a = 1
	//   b
	//   c
	//   d = 3
	//   f
	// )
	// fmt.Println(a, b, c, d, f)

	// var devide_1 float32 = 10.0 / 4
	// var devide_2 float32 = 10 / 4.001
	// fmt.Println(devide_1)
	// fmt.Println(devide_2)

	// array_1 := [3]int{1,3,4}
	// fmt.Println(array_1)

	// var users = [...]string{"Tom", "Alice", "Kate"}

	// for index, value := range users {
	// 	fmt.Println(index, value)
	// }

	// for i := 0; i < len(users); i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(users[i])
	// }

	// add(1, 2, 4, 6, 1)
	// var a = add([]int{5, 6, 7, 2, 3}...)
	// fmt.Println(a)
}
