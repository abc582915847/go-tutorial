package main

import (
	"fmt"
	"os"
)

func main() {

	//4
	person := Person{"hello",10}

	println("person-->",person)

	s1, s2 := getStatus(10, "hello")
	_, s3 := getStatus(10, "hello")
	if s2 == true {
		println("\n", s1, s2)
	}
	println("s3---",s3)
	println("\n")

	//3
	var result = add(10, 100)
	println("result----", result)
	println("\n")

	//2
	testAssignValue()

	//1
	helloWorld()

	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Printf("\nit is over ----%s", os.Args[1]);
}

func helloWorld() {
	fmt.Print("hello")
	println("\n")
	println("hello go world")
}

func testAssignValue() {
	var test string
	test, test1 := "test", 9000

	println(test1)

	var hello string = "hello"

	println(hello)
	println(test)

	var mHello = getStr()
	println(mHello)

	fmt.Printf("mhello----%s", mHello)

}

func getStr() string {
	return "\nmhello"
}

func add(val1, val2 int8) int8 {
	return val1 + val2
}

func getStatus(val int, s1 string) (string, bool) {
	return s1, true
}
