package main

import (
	"errors"
	"fmt"
	"math"
)

// **GoLang Creating Structs** //

type person struct {
	name string
	age  int
}

func main() {
	// **GoLang INITIALIZE VARIABLES** //

	fmt.Println("** Initialize Variables **")
	x := 5
	var y int = 7
	sum := x + y
	fmt.Println(sum)
	fmt.Println("------------")

	// **GoLang IF STATEMENTS** //

	fmt.Println("** If Statements **")
	if sum > 6 {
		fmt.Println("More than 6!")
		fmt.Println("------------")
	} else if sum < 3 {
		fmt.Println("Less than 3!")
	} else {
		fmt.Println("Yeet")
	}

	// **GoLang INITALIZE ARRAYS** //

	fmt.Println("** Initialize Arrays + Slices **")
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	A := [5]int{1, 2, 3, 4, 5}
	fmt.Println(A)

	b := []int{5, 4, 3, 2, 1} // Array "Slices" allow you to make appends to array lists
	b = append(b, 69)
	fmt.Println(b)
	fmt.Println("------------")

	// **GoLang INTITALIZE MAPS** (similar to dictionaries in Python) //
	fmt.Println("** Initialize Maps **")
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	delete(vertices, "square")
	fmt.Println(vertices)
	fmt.Println("------------")

	// **GoLang LOOPS** //
	fmt.Println("** Initialize Loops **")
	for i := 0; i < 5; i++ {
		fmt.Println(i, "FOR")
	}

	fmt.Println(" ")

	j := 0
	for j < 5 {
		fmt.Println(j, "WHILE")
		j++
	}
	fmt.Println("------------")

	// **GoLang Loops w/ Arrays + Maps** //
	fmt.Println("** Loops w/ Arrays + Maps **")
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	fmt.Println(" ")

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}
	fmt.Println("------------")

	// **GoLang Calling Functions** //
	fmt.Println("** Calling Functions **")
	result := add(2, 3)
	fmt.Println(result)

	fmt.Println(" ")

	r, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	fmt.Println("------------")

	// **GoLang Initializing Structs/Types** //
	fmt.Println("** Initializing Structs/Types **")
	p := person{name: "Noah", age: 21}
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println("------------")

	// **GoLang Pointers** //
	fmt.Println("** Pointers **")

	point := 7
	inc(&point) // Using "&" points to the specific memory address the variable is stored
	fmt.Println(point)

}

// **GoLang Initializing New Functions //
func add(x int, y int) int { // func -> name of function -> (arguments) -> return type -> {function logic}
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

// **GoLang Pointer Functions** //
func inc(x *int) {
	*x++
}
