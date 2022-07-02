package main

import "fmt"

func bye() {
	fmt.Println("Good bye !")
	callByOtherSourceFile("bye func")
}
