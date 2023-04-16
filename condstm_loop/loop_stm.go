package condstmloop

import "fmt"

func Loops() {
	//for loop
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\n", i)
	}

	//break statement - program terminates
	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d", i)
	}

	//continue statement - skip the current iteration of the for loop
	for i := 0; i <= 10; i++ {
		if i > 5 {
			continue
		}
		fmt.Printf("%d", i)
	}

	//nested loops
	for i := 5; i > 0; i-- {
		for j := 5; i > j; j-- {
			fmt.Println("nested for loop")
		}
	}

	//Labels - Labels can be used to break the outer loop from inside the inner for loop
outerloop:
	for i := 5; i > 0; i-- {
		for j := 5; i > j; j-- {
			fmt.Println("nested for loop")
			break outerloop
		}
	}

	//Initialisation, Conditional, Post are optional
	i := 0
	for i <= 10 {
		fmt.Printf("Initialisation, Conditional, Post are optional %d", i)
	}

	//multiple variables in for loop
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//Infinity loop
	for {
	}
}
