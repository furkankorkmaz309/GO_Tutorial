package main

import "fmt"

func main() {
	var integer int = 5
	//e.g. 5, 12, -4 etc.
	fmt.Println("integer :", integer)

	var float float64 = 5.2
	//e.g. 0.5, -3.7 etc.
	fmt.Println("float :", float)

	var string string = "Hello World"
	// Note: Write with double quotes or backticks. NOT SINGLE QUOTE
	//e.g. `Hi everyone'
	fmt.Println("string :", string)

	var boolean bool = true
	//true or false
	fmt.Println("boolean :", boolean)

	var uint uint = 5
	//e.g. 0, 10, 225 etc.
	// Note: non-negative numbers
	fmt.Println("uint :", uint)

	var int32 int32 = 2147483647
	// 32 bit signed integers between -2,147,483,648 and 2,147,483,647
	fmt.Println("int32 :", int32)

	var uint32 uint32 = 4294967295
	// 32 bit unsigned integers between 0 and 4,294,967,295
	fmt.Println("uint32 :", uint32)

}
