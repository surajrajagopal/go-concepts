package condstmloop

import "fmt"

func SwitchConditions() {
	//expression
	num := 1
	switch num {
	case 1:
		fmt.Println("")
	case 2:
		fmt.Println("")
	case 3:
		fmt.Println("")
	}

	//duplicate cases are not allowed
	// num1 := 0
	// switch num {
	// case 1:
	// 	fmt.Println("")
	// case 2:
	// 	fmt.Println("")
	// case 2:
	// 	fmt.Println("")
	// }

	//default case
	num1 := 0
	switch num1 {
	case 1:
		fmt.Println("")
	case 2:
		fmt.Println("")
	case 3:
		fmt.Println("")
	default:
		fmt.Println("")
	}

	//Multiple expressions in case
	character := "i"

	switch character {
	case "a", "e", "i":
		fmt.Println("Multiple expressions in case")
	default:
		fmt.Println("character not available")
	}

	//expressionless switch

	num2 := 10

	switch {
	case num2 > 10:
		fmt.Println("case 1")
	default:
		fmt.Println("default case ")
	}
}
