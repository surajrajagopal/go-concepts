package firstorderfunctions

import "fmt"

/*
higher order functions
https://go.dev/play/
*/

func Anonymous() {
	a := func() {
		fmt.Println("Anonymous function which assign to variable")
	}
	a()

	func() {
		fmt.Println("Anonymous function without pass arguments")
	}()

	func(name string) {
		fmt.Printf("My name is %s", name)
	}("Sean")
}
