package functionspackage

//function
func CustomFunction() string {
	return "Hello string"
}

//Multiple return value function
func MultipleReturnValues() (string, string) {
	return "Hello", "Worlds"
}

//Named return values
func NamedReturnValues() (a, b string) {
	a = "i am 'A' named return value"
	b = "i am 'A' named return value"
	return
}

//Blank Identifier - (_) identifier is used to discard the perimeter
