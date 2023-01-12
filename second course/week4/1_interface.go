package main

/* Concrete types Vs interface types
concrete type
- Specify the exact representation of the data and methods
- Complete method implementation is included

interface types
- specify some methods signatures
- implementations are abstracted

* NOTE: the interface is always mapped to a concrete type
*/

/* Interface values
it can be treated like other values
- Assigned to a variable
- Passed and returned

interface values have two components
Dynamic type: The concrete type which it is assigned to
Dynamic value: The value of the dynamic type
* NOTE: an interface value can have a nil dynamic value
* (it must have a dynamic type)
*/
import "fmt"

type Speaker interface{ speak() }

type Dog struct{ name string }
type Cat struct{ name string }

func (d Dog) speak() {
	fmt.Println(d.name)
}

//the method below shows the implementation of nil dynamic value
func (c *Cat) speak() {
	if c == nil {
		fmt.Println("Noise")
	} else {
		fmt.Println(c.name)
	}
}

func main() {
	d1 := Dog{"Brian"}
	var s1 Speaker
	s1 = d1
	// *dynamic type -> dog, dynamic value -> d1
	d1.speak()
	s1.speak()

	// Implementation of nil dynamic value
	var s2 Speaker
	var c1 *Cat
	s2 = c1
	s2.speak()
}
