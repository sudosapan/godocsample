package godocsample

import "fmt"

//Package level example
func Example() {
	fmt.Println("Package level example")
}

//Package level example
func Example_another() {
	fmt.Println("Package level another example")
}

//Testable example of Add funtion
func ExampleAdd() {
	a := 10
	b := 20
	fmt.Println(Add(a, b))
	//Output: 30
}

// example of Add funtion, does not add to tests
func ExampleAdd_nontestable() {
	a := 10
	b := 20
	fmt.Println(Add(a, b))
}

// demonstrates using the person struct, and is a testable example too.
func ExamplePerson() {
	person := &Person{Name: "sudosapan"}
	fmt.Println(person.Name)
	//Output: sudosapan
}

// example of Person method GetAge
func ExamplePerson_GetAge() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}

// example of Person method GetAge
func ExamplePerson_GetAge_test1() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}

// example of Person method GetAge
func ExamplePerson_GetAge_test2() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}

// This will not show up in the example section due to more than accepted "_".
//But this will add to tests being executable example
func ExamplePerson_GetAge_test_3() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}
