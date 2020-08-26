package main

import "fmt"

func main() {
	a := 1
	fmt.Printf(helloWorld("applyboy"))
	fmt.Println("learn go in one day")

	if a >= 1 {
		fmt.Println("a >= 1")
	}
}

func helloWorld(userName string) string {
	return fmt.Sprintf("Hi, %s", userName)
}