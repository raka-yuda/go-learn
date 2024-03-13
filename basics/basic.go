package main

import (
	"errors"
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

type Address struct {
	City, Province, Country string
}

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "1001" {
		return NotFoundError
	}

	return nil
}

func main() {
	// Print
	fmt.Print("Hello")
	fmt.Println("Hello, Guysss")
	fmt.Printf("Hello, %s", "Guysss \n")

	// Var
	// var a, b, c, d int = 1, 3, 5, 7

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	const PI = 3.14
	fmt.Println(PI)

	// Print output
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)

	// Array
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// Slice
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Operator
	var (
		sum1 = 100 + 50    // 150 (100 + 50)
		sum2 = sum1 + 250  // 400 (150 + 250)
		sum3 = sum2 + sum2 // 800 (400 + 400)
	)
	fmt.Println(sum3)

	var x = 10
	x += 5
	fmt.Println(x)

	var x1 = 5
	var y2 = 3
	fmt.Println(x1 > y2)

	// Condition
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	// Switch
	day := 5

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// Loops
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

	fruits2 := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits2 {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// Function
	myMessage()

	// Recursion
	testcount(1)

	// Struct
	var pers1 Person

	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Maps
	var a1 = make(map[string]string) // The map is empty now
	a1["brand"] = "Ford"
	a1["model"] = "Mustang"
	a1["year"] = "1964"

	fmt.Printf("a\t%v\n", a1)

	// Comment
	// This is single comment

	/*
		This is
		multiline comment
	*/

	// Pointer
	//address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1 // copy value

	//address2.City = "Bandung"
	//fmt.Println(address1) // tidak berubah
	//fmt.Println(address2) // berubah menjadi Bandung

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung

	// Custom Error
	err := GetById("1001")
	if err != nil {
		// Used 'Is' function to check which error
		if errors.Is(err, ValidationError) {
			fmt.Println(ValidationError.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println(NotFoundError.Error())
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
