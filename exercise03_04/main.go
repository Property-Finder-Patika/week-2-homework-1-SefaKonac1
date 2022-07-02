package main

import "runtime"

func main() {

	greet()
	bye()
	getGoVersion()
}

//unified function call
func callByOtherSourceFile(caller string) {
	println("scope exercise !" + caller + "\n")
}

func getGoVersion() {
	println(runtime.Version())
}
