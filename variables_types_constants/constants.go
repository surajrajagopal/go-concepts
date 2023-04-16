package variablestypesconstants

import (
	"fmt"
	"math"
)

func Constants() {
	//Declaring a constant
	const a = 10
	fmt.Printf("%T", a)

	//Declaring a group of constants
	const (
		name    = "John"
		age     = 50
		country = "Canada"
	)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	// This is not allowed since b is a constant
	const b = 55 //allowed
	//   b = 89 //reassignment not allowed

	/*The value of a constant should be known at compile time.
	Hence it cannot be assigned to a value returned by a function call
	since the function call takes place at run time*/
	var a1 = math.Sqrt(4) //allowed
	//const b1 = math.Sqrt(4) //not allowed
	fmt.Println(a1)
}
