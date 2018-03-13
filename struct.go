/*
    Ken Bailey
    3/12/18

    All about structs in Go

*/

package main

import (
    "fmt"
    "reflect"
)

// Go has several built in types that most people are familiar with.
// Go also allows you to create more complex types based on these foundational types
type Person struct {
    age         int
    name        string
    weight      float32
    isJamaican  bool
}

// Go allows us to create a type alias like this.
// These two types are identical b/c their underlying type
//  literals are structurally equivalent.
type OtherPerson Person

func main() {
    // We can instantiate a Person object
    //  and initialize its properties on the same line in the
    //  order they are declared in the struct.
    ken := Person{25, "ken", 175.5, true}
    fmt.Println(ken.age)

    // We can also Instantiate a Person object
    //  and initialize its properties one by one. Upon instantiation 
    //  of an object without assigning its properties values, each 
    //  property is initialized with a default value
    karla := Person{}
    karla.name = "karla"
    fmt.Printf("%s: %d\n", karla.name, karla.age)
    karla.age = 22
    fmt.Printf("%s: %d\n", karla.name, karla.age)

    //  If we wanted to cast a one type into another we use "TYPE()".
    //   We use reflection to get the the variable's type.
    notKen := OtherPerson{90, "notken", 300.3, false}
    otherKen := Person(notKen)
    fmt.Println("notKen type: ", reflect.TypeOf(notKen))
    fmt.Println("otherKen type: ", reflect.TypeOf(otherKen))

    // Even though types Person and OtherPerson are stucturally similar,
    //  we cant directly assign them to each other. Statements below are
    //  illegal.
    //
    //  var newKen Person
    //  newKen = notKen

}