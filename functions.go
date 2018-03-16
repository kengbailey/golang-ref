/*
    Ken Bailey
    3/7/18

    All about Go functions

    Notes:
    - Functions in Go can have named return parameters. Upon entering the function
        they are initialized with zero values for their types and if the function 
        executes a return statement without any parameters the named parameters 
        are used as return values
    - Go's defer statement schedules a function call(the function being deferred) to 
        be run immediately before the calling function executing the defer returns
    - Any parameters passed to a deferred function are evaluated when the defer 
        executes and not when the call executes!
    - Go allows for anonymous functions, also known as closures. These can be 
        created within other functions and must be called with parentheses after
        they are defined.

*/

package main

import (
    "strings"
    "fmt"
    "strconv"
)

// Declare a function with two input parameters.
// Variable of the same type can be declared at the same time; separated by commas.
// Here we are just concatenating two strings
func StringTogether(a, b string) string {
    return a+b
}

// Declare a function that returns multiple values.
// You can return as many values as needed.
// Here we are testing to see if a string contains "ken" and returning a bool and 
//  the len of the string.
func ContainsKen(s string) (i int, ans bool) {
    ans = strings.Contains(s, "ken")
    i = len(s)
    return
}

// Declare a function that returns a function.
// Function being returned must be named 'func' and have same # of params.
// Params to function being returned don't have to be named until return statement.
func ReturnFunc(s string) func(string) string {
    return func(x string) string {
        return s+x
    }
}

func main() {
    // Call StringTogether function
    fmt.Println(StringTogether("Ken Bailey", " is a Jamaican"))

    // Call ContainsKen function
    length, containsKen := ContainsKen("token")
    if(containsKen) {
        fmt.Println("token: " + strconv.Itoa(length))
    }

    // Call ContainsKen function
    // To toss a value returned from a function we use "_", the blank identifier
    _, contains := ContainsKen("awoken")
    fmt.Println(contains)

    // Call to ReturnFunc where a function is returned and stored in f
    f := ReturnFunc("Ken wants to go back to Jamaica")
    newString := f(" today!")
    fmt.Println(newString)

    // Here we create an anonymous function, also known as a closure in golang.
    // It must be called after it is defined.
    // This actually will be run, but it's printing won't be seen b/c we aren't 
    //  aren't waiting for its execution to finish! 
    // The goroutine is started and immediately stopped b/c the main function 
    //  stops immediately after we kick off the goroutine.
    go func(s string) {
        fmt.Printf("%s, one of these days\n", s)
    }("Ken needs to go to Jamaica")
}
