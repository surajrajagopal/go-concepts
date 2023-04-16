package pointers

import "fmt"

func DeclarePointers() {
	b := 100
	// a is pointing to b
	// * - operator use to accessing the value of the variable to which the pointer points
	// & - operator use to get address of variable
	var a *int = &b
	fmt.Printf("a type %T\n", a)
	fmt.Println("addrres of b\n", a)
	fmt.Printf("read pointer value %d", *a)
}

func PointerWithNew() {
	ptr := new(string)
	fmt.Printf("Size value is %s, type is %T, address is %v\n", *ptr, ptr, ptr)

	*ptr = "new function"
	fmt.Println("New size value is", *ptr)
}

func DeferencePointer() {
	a := 10
	b := &a
	fmt.Println("address of b is", b)
	fmt.Println("value of b is", *b)
}

func PointerTofunc(num *int) *int {
	*num += 1
	return num
}

//It is perfectly legal for a function to return a pointer of a local variable ex: i var
func ReturnPointerFromFunc() *int {
	i := 10
	return &i
}

func ArrayPointer(sli *[3]int) *[3]int {
	sli[0] = 10
	return sli
}

func SlicePointer(sli []int) []int {
	sli[0] = 10
	return sli
}
