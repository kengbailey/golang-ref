/*
    Ken Bailey
    3/7/18

    All about go functons
 
*/

package main

import (
    "strings"
    "fmt"
    "strconv"
)

// Declare a function with two input parameters.
//  Variable of the same type can be declared at the same time; separated by commas.
//  Here we are just concatenating two strings
func StringTogether(a, b string) string {
    return a+b
}

// Declare a function that returns multiple values
//  You can return as many values as needed.
//  Here were are testing to see if a string contains "ken" and returning a bool and the len of the string
func ContainsKen(s string) (int, bool) {
    ans := strings.Contains(s, "ken")
    return len(s), ans
}

// Declare a function that returns a function
//  Function being returned must be named 'func' and have same # of params
//  Params to function being returned don't have to be named until return statement
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
    //  To toss a value returned from a function we use the _ character
    _, contains := ContainsKen("awoken")
    fmt.Println(contains)

    // Call to ReturnFunc
    f := ReturnFunc("Ken wants to go back to Jamaica")
    newString := f(" today!")
    fmt.Println(newString)

}
