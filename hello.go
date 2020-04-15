package main

import "fmt"

func main() {
	fmt.Println("Before")
	foo()
	fmt.Println("After")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("In foo function")
}

func bar() {
	fmt.Println("We exited control flow")
}
