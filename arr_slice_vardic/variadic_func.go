package arrslicevardic

import "fmt"

func VaradicFunc(num ...int) {
	n := num[0]
	switch n { //expression
	case 1:
		fmt.Println(n)
	case 2:
		fmt.Println(n)
	case 3:
		fmt.Println(n)
	case 4:
		fmt.Println(n)
	default:
		fmt.Print("value not found")
	}
}
