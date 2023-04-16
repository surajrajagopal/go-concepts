package variablestypesconstants

import "fmt"

func Variable() {
	//Declaring a single variable
	var age int
	fmt.Printf("%d", age)

	//Declaring a variable with an initial value
	var age1 int = 29
	fmt.Printf("%d", age1)

	//Type inference - If a variable has an initial value, Go will automatically be able to infer the type of that variable using that initial value
	var age2 = 29 // type will be inferred
	fmt.Println("My age is", age2)

	//Multiple variable declaration
	var width, height int
	fmt.Print(width, height)

	//Multiple variable declaration with intialization
	var width1, height1 int = 100, 50 //declaring multiple variables
	fmt.Print(width1, height1)

	//Multiple variable declaration with type inference
	var width2, height2 = 100, 50
	fmt.Print(width2, height2)

	//variables belonging to different types in a single statement
	var (
		num1  = 10
		str   = "hello"
		float = 5.4
	)
	fmt.Println(num1, str, float)

	//Short hand declaration
	num2 := 100
	fmt.Println(num2)

	//multiple variables in a single line using short hand syntax
	num3, str1 := 10, "hello"
	fmt.Println(num3, str1)

	//Short hand syntax can only be used when at least one of the variables on the left side of := is newly declared.
}
