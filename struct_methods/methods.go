package structmethods

import "fmt"

type PersonalDetails struct {
	Name    string
	PhoneNo int
	Address string
}

func (p PersonalDetails) Details() {
	res := fmt.Sprintf("my name is %s, no : %d, address %s", p.Name, p.PhoneNo, p.Address)
	fmt.Println(res)
}

/*Pointer Receivers vs Value Receivers - The difference between value and pointer receiver is,
changes made inside a method with a pointer receiver is visible to the caller
whereas this is not the case in value receiver.*/

/*
Pointer Receiver vs Value Receiver
Example link : https://go.dev/play/p/1tuBmFtLJ5R

-Methods of anonymous struct fields
Example link : https://go.dev/play/p/QXeJpOrwvUZ

-Value argument in functions vs Value receiver in method
Example link : https://go.dev/play/p/MfihRLCki4n

Pointer argument in functions vs Pointer receiver in method
Example link : https://go.dev/play/p/WpPqL6Ctr22

Methods with non-struct type
Example link : https://go.dev/play/p/xUtPBM3hEJ6
*/
