package main

import "fmt"

func main() {
	fmt.Println("this is new test for go!!!")

	foo()

	for i:=0; i<5; i++ {
		fmt.Println("this is the loop", i)
	}
}

func foo()  {
	fmt.Println("I am in foo")
}

