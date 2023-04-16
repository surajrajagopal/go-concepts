package mapstring

import (
	"fmt"
	"unicode/utf8"
)

//Note : len(s) is used to find the number of bytes in the string.
func LenAndRuneCount() {
	word := "Señor"
	fmt.Printf("Length of %d\n", utf8.RuneCountInString(word))
	fmt.Printf("Number of bytes %d\n", len(word))
}

func DeclareString() string {
	str := "Constant string value"
	return str
}

func PrintBytes(s string) {
	//Print HexaDecimal Value of input, %x format specifier of HexaDecimal
	fmt.Print("Bytes : ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func PrintChars(s string) {
	//Print character Value of input, %c format specifier of character
	fmt.Print("Characters : ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

//Rune represents a Unicode code point (int32) in Go. It doesn't matter how many bytes the code point occupies.
func PrintRune(s string) {
	//Print character Value of input, %c format specifier of character
	r := []rune(s)
	fmt.Print("Characters : ")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c ", s[i])
	}
}

//AccessRuneByForRange
func AccessRuneByForRange(s string) {
	fmt.Print("Rune : \n")
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

//CreateStringFromByteSlice
func CreateStringFromByteSlice() {
	//Decimal
	byteSlice := []byte{67, 97, 102, 195, 169}
	//hexadecimal
	byteSlice1 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}

	str := string(byteSlice)
	str1 := string(byteSlice1)
	fmt.Println("Decimal Values to String", str)
	fmt.Println("HexaDecimal Values to String", str1)
}

//CreateStringFromRuneSlice
func CreateStringFromRuneSlice() {
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str := string(runeSlice)
	fmt.Println("HexaDecimal Values to String", str)
}

//== operator is used to compare two strings for equality
func CompareStrings(str1, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal", str1, str2)
}

//String Concatenation
// 1. + - operator
// 2. Inbuilt func - fmt.Sprintf()
func StringConcatenation(str1, str2 string) {
	//way 1
	s := str1 + " " + str2
	fmt.Printf("string concatenation way 1 %s", s)

	//way 2
	result := fmt.Sprintf("string concatenation way 2 %s %s", str1, str2)
	fmt.Println("string concatenation way 2", result)

}

//String are immutable - Once a string is created it's not possible to change it
func immutate(s string) string {
	//	s[0] = 'a' //any valid unicode character within single quote is a rune
	return s
}

// Convert to rune and mutate string type
func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
