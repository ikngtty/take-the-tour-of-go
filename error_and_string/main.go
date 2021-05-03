package main

import "fmt"

type MyType1 struct {
	Value int
}

func (str MyType1) String() string {
	return "I am a stringer!"
}
func (err MyType1) Error() string {
	return "I am an error!"
}

type MyType2 struct {
	Value int
}

func (str MyType2) String() string {
	return "I am a stringer!"
}

type MyType3 struct {
	Value int
}

func (err MyType3) Error() string {
	return "I am an error!"
}

func main() {
	mytype1 := MyType1{42}
	mytype2 := MyType2{42}
	mytype3 := MyType3{42}

	fmt.Println(mytype1) // -> I am an error!
	fmt.Println(mytype2) // -> I am a stringer!
	fmt.Println(mytype3) // -> I am an error!

	// CONCLUSION: As error > As fmt.Stringer > As an other type
}
