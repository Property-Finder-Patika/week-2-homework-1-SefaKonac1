package main

import "fmt"

func main() {
	myName := "sefa"
	myFriendsName := "erdem"
	printNamesWithPrint(myName, myFriendsName)
	printNamesWithPrintln(myName, myFriendsName)
	printSeperatedValues(myName, myFriendsName)
}

func printNamesWithPrint(name1 string, name2 string) {
	fmt.Println("println :" + name1)
	fmt.Println("println : " + name2)
}

func printNamesWithPrintln(name1 string, name2 string) {
	print("print :" + name1)
	print("print : " + name2)
}

func printSeperatedValues(name1 string, name2 string) {
	fmt.Println("println with seperated by coma:", name1, name2)
	print("print with seperated by comma:", name1, name2)
}
