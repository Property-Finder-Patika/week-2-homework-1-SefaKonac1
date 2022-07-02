package main

import "fmt"

func greet() {
	fmt.Println("hello ! I'm greet page.")
	callByOtherSourceFile("greet func")
}
