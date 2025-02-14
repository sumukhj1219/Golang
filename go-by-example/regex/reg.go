package main

import (
	"fmt"
	"regexp"
)

func main() {
	te := "b14f6e1a-78c3-4b8d-a3f1-6b2a45eac9d9"
	re := regexp.MustCompile(`^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`)
	if !re.MatchString(te) {
		fmt.Println("Invalid uuid ❌")
	} else {
		fmt.Println("Matched ✅")
	}
	// mm := re.FindAllString(te, -1)
	// id := re.FindAllStringIndex(te, -1)
	// up := re.ReplaceAllStringFunc(te, strings.ToUpper)
	// fmt.Println(up, te)
}
