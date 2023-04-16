package variablestypesconstants

import "fmt"

func Types() {
	//bool
	var a = true
	var b = false
	fmt.Println(a && b) //both condition has to true

	var a1 = true
	var b1 = false
	fmt.Println(a1 && b1) //anyone of condition has to true

	//int - signed and unsigned types
	var num int = 10
	fmt.Println(num)

	var num1 uint
	fmt.Println(num1)

	//float - float32 and float64
	var f float32 = 8.9
	fmt.Println(f)

	var f1 float64 = 8.9
	fmt.Println(f1)

	//Complex types

	//byte - is an alias of uint8

	//rune is an alias of int32

	//Strings are a collection of bytes in Go.

	//Type Conversion
	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)

	i1 := 10
	var j1 float64 = float64(i1) //this statement will not work without explicit conversion
	fmt.Println("j", j1)
}
