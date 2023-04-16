package condstmloop

import "fmt"

func ConditionalStatements() {
	//if statement
	if true {
		fmt.Println("true")
	}

	//if-else statement
	if false {
		fmt.Println("if-else statement")
	} else {
		fmt.Println("if-else statement")
	}

	//if-elseif-else
	if true {
		fmt.Println("if-elseif-else")
	} else if false {
		fmt.Println("if-elseif-else")
	} else {
		fmt.Println("if-elseif-else")
	}

	//if with assignment statement
	if num := 10; num > 1 {
		fmt.Println("if with assignment statement")
	}

	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")
}
