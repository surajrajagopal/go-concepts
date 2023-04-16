package structmethods

// Named struct type
type Details struct {
	FirstName string
	PhoneNo   int
}

type Address struct {
	Street, LandMark string
	Pincode          int
}

//Accessing individual fields of a struct
//The dot . operator is used to access the individual fields of a struct.

/*Zero value of a struct - When a struct is defined and it is not explicitly initialized with any value,
the fields of the struct are assigned their zero values by default.
It is also possible to specify values for some fields and ignore the rest.
In this case, the ignored fields are assigned zero values.
*/

// Anonymous fields
type AnonyFields struct {
	string
	int
}

// Nested structs - It is possible that a struct contains a field which in turn is a struct. These kinds of structs are called nested structs
type NestedStruct struct {
	address Address
}

// Promoted field-
type Struct struct {
	Address
}

// Exported structs and fields
type Address1 struct { //exported struct
	Street   string //exported field
	LandMark string //exported field
	Pincode  int    //unexported field
}

/*Structs Equality - 1. Structs are value types and are comparable if each of their fields are comparable.
Two struct variables are considered equal if their corresponding fields are equal.
2. Struct variables are not comparable if they contain fields that are not comparable
*/

/*
//Anonymous struct
emp3 := struct {
	name string
}{
	name: "Hello",
}
fmt.Println(emp3)

//slice of Anonymous struct
emp4 := []struct {
	name string
}{
	{
		name: "Hello",
	},
}
fmt.Println(emp4)
*/

/*
	// with specifing field name of struct type
	det := structmethods.Details{
		FirstName: "Hello",
		PhoneNo:   1234,
	}

	// without specifing field name of struct type
	addr := structmethods.Address{
		"area",
		"world",
		123456,
	}
	fmt.Println("Details struct", det)
	fmt.Println("Address struct without specifing", addr)

	// Pointer to a struct
	addr1 := &structmethods.Address{
		Street:   "area",
		LandMark: "YP",
		Pincode:  123456,
	}
	fmt.Println("Address struct without specifing", addr1)
*/
