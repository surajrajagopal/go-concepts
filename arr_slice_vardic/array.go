package arrslicevardic

import "fmt"

func Array() {
	//Note : The size of the array is a part of the typey
	// a1 := [3]int{5, 78, 8}
	// var b1 [5]int
	//	b1 = a1 //not possible since [3]int and [5]int are distinct types

	// with var declaring array
	var arr [3]int
	arr[0] = 10
	arr[1] = 11
	fmt.Println("this is array", arr)

	//short hand declaration
	ar := [3]int{12}
	fmt.Println("this is short hand declaration", ar)

	//use ... operator to declare array
	arr1 := [...]int{1, 6, 7}
	fmt.Println(arr1)
	//find the length of arr
	fmt.Println(len(arr1))

	// for loops
	for i := 0; i < len(arr1); i++ { //looping from 0 to the length of the array
		fmt.Printf("%d th element of a is %d\n", i, arr1[i])
	}

	//for range loop
	for i, v := range arr1 {
		fmt.Printf("index %d and value %d", i, v)
	}

	//Multi-Dimensional Array with var declaration
	var mulArr [3][2]int
	mulArr[0][0] = 10
	mulArr[0][1] = 10
	fmt.Println(mulArr)

	//Multi-Dimensional Array withshort hand declaration
	mulArray := [3][2]string{
		{"hello", "world"},
		{"hello", "world"},
		{"hello", "world"},
	}
	fmt.Println(mulArray)
}
