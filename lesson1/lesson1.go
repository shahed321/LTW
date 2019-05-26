package lesson1

import "fmt"

func Lesson1() {
	var output string
	fmt.Println("Please Enter Your name: ")
	fmt.Scanln(&output)
	fmt.Printf("Hello %s \n", output)
}
