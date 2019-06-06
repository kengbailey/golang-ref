package main

import "fmt"

func main() {
	test := permutations("ken")
	for _, t := range test {
		fmt.Println(t)
	}
	fmt.Println(len(test))
}

// permutations finds all permutations of a given string
func permutations(str string) (perms []string) {
	if len(str) == 1 { // base case
		perms = append(perms, str)
	} else {
		for i := 0; i < len(str); i++ {
			// find all permuations of string minus character @ current index(i)
			permutationsMinusI := permutations(string(str[0:i]) + string(str[i+1:]))
			// assemble permuatations of current string with character @ current index(i) in front
			for _, p := range permutationsMinusI {
				perms = append(perms, string(string(str[i])+p))
			}
		}
	}
	return
}
