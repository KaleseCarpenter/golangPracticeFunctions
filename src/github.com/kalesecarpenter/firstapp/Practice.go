package main

import "fmt"

/*func main() {

	// this is saying I am decaring a variable i, that is of TYPE integer
	//		var i int //when you want to decalre a variable but youore not ready to initialize it yet, or local variable scope
	//		i = 42 //
	var j int = 27
	//		k := 99
	// i := 42 //this is the same as var i int = 42
	// i = 27 // this statement can only do one thing at a time

	//allows you to print out a formatting string
	fmt.Printf("%v, %T", j, j)
}*/

/*func main() {

	var j float32 = 27

	//allows you to print out a formatting string
	fmt.Printf("%v, %T", j, j)
}*/

/*func main() {

	k := 99.
	// k := 99

	//allows you to print out a formatting string
	fmt.Printf("%v, %T", k, k)
}*/

/*func main() {

	var i float32 = 42
	// k := 99

	//allows you to print out a formatting string
	fmt.Printf("%v, %T", i, i)
}*/

/*func main() {
	// keep the code clean by grouping variable names
	var (
		actorName    string = "Elisabeth Sladen"
		companion    string = "Sarah Jane Smith"
		doctorNumber int    = 3
		season       int    = 11
	)

	var (
		counter int = 0
	)

	fmt.Printf("%v, %T", i, i)
}*/

//Variables
/*var i int = 27

func main() {
	// This is to print out the Global variable first
	fmt.Println(i)
	// Shadowing - the variable with the inner most scope takes precedent
	var i int = 42
	fmt.Println(i)
}*/

/*func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	// conversion functions have to be explicitly converted when changing TYPES because floating point numbers can be decimals
	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j) // \n means - line feed or newline
}*/

// CONVERT AN INTEGER TO A STRING
/*func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	// conversion functions have to be explicitly converted when changing TYPES because floating point numbers can be decimals
	var j string
	j = strconv.Itoa(i)          // means String conversion I to an ASCII string -a string of data as ASCII characters that can be interpreted and displayed as readable plain text for people and as data for computers.
	fmt.Printf("%v, %T\n", j, j) // \n means - line feed or newline
}*/

// BOOLEANS

/*func main() {
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

}*/

// INTEGERS
/*func main() {
	n := 42

	fmt.Printf("%v, %T\n", n, n)

}*/

/*func main() {
	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // this prints out 3 instead of 3 remainder 1 because when you divide an integer by an integer, you have to get an integer as the result
	fmt.Println(a % b)

}*/

// ***** DATA CONVERSION INTEGERS

/*func main() {
	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b)) // you have to put the int in front of the variable for th type conversion
}*/

// BIT Shifting

/*func main() {
	var a := 8 // 2^3
	fmt.Println(a << 3 ) // shift to the left   2^3 * 2^3 = 2^6
	fmt.Println(a >> 3 ) // shift to the right 2^3 / 2^3 = 2^0
}*/

// BIT Operators --

func main() {
	a := 10 // 1010 -- Binary Representation
	b := 3  // 0011 --

	fmt.Println(a & b)  // 0010 -- & is looking for what bits are set in first and second number
	fmt.Println(a | b)  // 1011 -- if 1 or the other is set
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100
}
