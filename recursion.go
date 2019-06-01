package main

import "fmt"

func main() {
	test := permutations("school")
	for _, t := range test {
		fmt.Println(t)
	}
	fmt.Println(len(test))
}

// permutations finds all permutations of a given string
func permutations(str string) []string {
	var perms []string
	if len(str) == 1 { // base case #1
		perms = append(perms, str)
	} else if len(str) == 2 { // base case #2
		perms = append(perms, str)
		perms = append(perms, string(str[1])+string(str[0]))
	} else {
		for i := 0; i < len(str); i++ {
			if i == 0 {
				outPerms := permutations(string(str[1:]))
				for _, p := range outPerms {
					perms = append(perms, string(string(str[0])+p))
				}
			} else {
				outPerms := permutations(string(str[0:i]) + string(str[i+1:]))
				for _, p := range outPerms {
					perms = append(perms, string(string(str[i])+p))
				}
			}
		}
	}
	return perms
}
