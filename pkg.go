//Package godocsample Sample package showing different ways go code can be documented.
package godocsample

//Add adds two numbers, and returns the sum
//
// call with int values
//   fmt.Println(Add(1,2))
func Add(a int, b int) int {
	return a + b
}

//Person denotes a person with a name
type Person struct {
	//Name of the person in question
	Name string
	// not exported, won't show in godoc
	birthyear int
}

//GetAge method on Person to get the age
func (p *Person) GetAge() int {
	return 2021 - p.birthyear
}
