package main

import (
	"fmt"
	"math"
)

/*this file includes source codes of

    - An exercise for arithmetic operations
    - An exercise for conversions among basic types
    - An exercise for conversions among string and basic types using strconv package
    - An exercise for iota
    - An exercise for scope
.*/

type (
	FinalMessage   string
	InitialMessage string
	Message        string
)

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func main() {

	IntegerToFloatConversion()
	isOdd(25)              //prints odd and returns true
	isOdd(16546548745413)  //prints odd and returns true
	isOdd(165465487454130) //prints even and returns false

	testValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	shiftingCoefficient := 2

	printSquareOfArray(testValues) //prints 0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
	printSqrtOfArray(testValues)   // printts sqrt of arrays elements value

	printRightShiftValueOfArray(testValues, shiftingCoefficient)
	printLeftShiftValueOfArray(testValues, shiftingCoefficient)

	big64Values := []int64{2, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 65536, 131072, 262144, 524288, 1048576}

	integerConversionObservers(big64Values)

	stringConversionObservers()
	iotaExample()
	scopeExercise()

}

func IntegerToFloatConversion() {
	x := 75
	y := 22

	result := float32(x) / float32(y)

	fmt.Printf("Integer to Float Conversion Example Result: {%v}", result)

}

func isOdd(n int) bool {

	if n%2 == 0 {
		fmt.Printf("%d is even", n)
		return false
	} else {
		fmt.Printf("%d is odd", n)
		return true
	}

}

func printSquareOfArray(values []int) {
	for el := range values {

		el *= 2
		fmt.Println(el)
	}
}

func printSqrtOfArray(values []int) {

	for el := range values {
		sqrtOfEl := math.Sqrt(float64(el))
		fmt.Println(sqrtOfEl)
	}
}

func printRightShiftValueOfArray(values []int, shiftingCoefficient int) {

	for el := range values {
		fmt.Printf("> {%d >> %d = %d} \n ", el, shiftingCoefficient, el>>shiftingCoefficient)
	}

}

func printLeftShiftValueOfArray(values []int, shiftingCoefficient int) {

	for el := range values {
		fmt.Printf("> {%d << %d = %d} \n ", el, shiftingCoefficient, el<<shiftingCoefficient)
	}

}

func integerConversionObservers(big64values []int64) {

	//converts 64 bits integer to 32, 8 integers
	for _, value64 := range big64values {
		value32 := int32(value64)
		value16 := int16(value64)
		value8 := int8(value64)
		fmt.Printf("> convert from int64 : %d  to int32: %d and to int16 : %d and to int8: %d \n", value64, value32, value16, value8)
	}

}

func stringConversionObservers() {

	var initial InitialMessage = "merhaba!"
	var end FinalMessage = "hoscakal!"

	fmt.Println(Message(initial) + Message(end))
}

func iotaExample() {

	fmt.Printf("The value of Red    is %v\n", Red)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Yellow is %v\n", Yellow)
	fmt.Printf("The value of Green  is %v\n", Green)
	fmt.Printf("The value of Blue   is %v\n", Blue)
	fmt.Printf("The value of Indigo is %v\n", Indigo)
	fmt.Printf("The value of Violet is %v\n", Violet)
}

/*the following code snipped was partly taken by book and implemented my own words*/
func scopeExercise() {

	x := "hello!" // outer x declaration

	for i := 0; i < len(x); i++ {
		x := x[i]     //loop x declaration points different memory location
		if x != '!' { //checks x defined into loop scope
			x := x + 65 // one more x declaration inside condititonal branch
			//on the other side, part of expression x is a value which defined
			//previous loop scope.
			fmt.Printf("%c", x)
		}
	}

}
