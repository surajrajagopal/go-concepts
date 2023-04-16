package structmethods

// call by value (without pointer) (changes will not reflect to caller)
type calculator struct {
	num1, num2 int
}

func (c calculator) SetValue(a, b int) (int, int) {
	c.num1 = a
	c.num2 = b
	return c.num1, c.num2
}

func (c calculator) Add() int {
	sub := c.num1 + c.num2
	return sub
}

func (c calculator) Sub() int {
	sub := c.num1 - c.num2
	return sub
}

/*
func main() {
	c := calculator{}
	fmt.Println(c.SetValue(4, 6))
	fmt.Println("reflected to caller", c)
	fmt.Println("Add values", c.Add())
	fmt.Println("Sub values", c.Sub())
}
output : 4 6
reflected to caller {0 0}
Add values 0
Sub values 0
*/
