# Go Documentation Samples

This repository contains examples for documenting go code and result of what it ends up in the go doc.


## Basic
### New Line
Give one blank comment before the text to go to new line
#### Comment
```go
// These will be
// shown in one line
//
// This starts in new line.
package godocsample
```

#### Go doc output:
![newline](/assets/newline.png)

### Heading
Text on newline starting with capital letter and not ending in full stop is rendered as a heading.

#### Comment
```go
//Package godocsample Sample package showing different ways go code can be documented.
//
// Heading
//
//This is under the heading.
//
package godocsample
```

#### Go doc output:
![heading](/assets/heading.png)


### Code
For getting comment getting rendered as code in the documentation, add more than 2 spaces between // and the text.

#### Comment
```go
// call with int values
//   fmt.Println(Add(1,2))
package godocsample
```

#### Go doc output:
![code](/assets/code.png)


## Package Documentation

Comments preceding the package directive in the go file go to the package documentation. 

There should be no (vertical) gap between the comment and package directive.
#### Comment
```go
// Sample package showing different ways go code can be documented.
package godocsample
```

#### Go doc output:
![package](/assets/package.png)

## Examples
You can include examples with optional output values in the code. Examples with output values specified will be executed and verified when the tests are run.

Examples are read from the _test file, so they need to be written there and linked to the owner with naming convention.
There should be not arguments or return of the Example function

Examples and the name convention is as 
| Level | Naming | Sample |
|-------|---------|----|
|Package|Example| ```func Example()``` OR ```func Example_<MorDetail>()```
|Function|Example\<FunctionName\>| ```func ExampleAdd()``` OR ```func ExampleAdd_<MoreDetail>()```|
|Struct|Example\<StructName\>| ```func ExamplePerson()``` OR ```func ExamplePerson_<MoreDetail>()```|
|Struct Method|Example\<StructName\>_\<MethodName\>| ```func ExamplePerson_GetAge()``` OR ```func ExamplePerson_GetAge_<MoreDetail>()```|


### Package Level example
In ![pkg_test.go](pkg_test.go)

#### Comment
```go
//Package level example
func Example() {
	fmt.Println("Package level example")
}
```

#### Go doc output:
![packageexample](/assets/packageexample.png)


### Package example another
In ![pkg_test.go](pkg_test.go)

#### Comment
```go
//Package level example
func Example() {
	fmt.Println("Package level example another")
}
```

#### Go doc output:
![packageexampleanother](/assets/packageexampleanother.png)

### Function Level Example
In ![pkg_test.go](pkg_test.go). 

#### Comment
```go
//Testable example of Add funtion
func ExampleAdd() {
	a := 10
	b := 20
	fmt.Println(Add(a, b))
	//Output: 30
}
```

Notice this example contains //Output: 30, this is the expected value. This will be displayed as output in documentation and will also be used to verify the result when tests are run.

Verify for yourself with 

```bash
go test
```

#### Go doc output:
![functionexample](/assets/functionexample.png)

If the //Output: tag is skipped then, documentation does not contain output and this example function is not part of the code tests.

### Struct Level Example
In ![pkg_test.go](pkg_test.go). 

#### Comment
```go
// demonstrates using the person struct, and is a testable example too.
func ExamplePerson() {
	person := &Person{Name: "sudosapan"}
	fmt.Println(person.Name)
	//Output: sudosapan
}
```

#### Go doc output:
![strcutexample](/assets/strcutexample.png)


### Struct Method Level Example
In ![pkg_test.go](pkg_test.go). 

#### Comment
```go
// example of Person method GetAge
func ExamplePerson_GetAge() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}
```

#### Go doc output:
![structmethod.png](/assets/structmethod.png)

There can be multiple example for struct methods

### Struct Method Level Example Mulitple
Like all others there can be multiple examples for one type. This can be acvieved by appending _ and detail to the Example. 

For example s

#### Comment
```go
// example of Person method GetAge
func ExamplePerson_GetAge_test1() {
	person := &Person{Name: "sudosapan", birthyear: 2020}
	fmt.Println(person.GetAge())
	//Output: 1
}
```
results in

#### Go doc output:
![strucmulti.png](/assets/strucmulti.png)



# Happy Documenting








