package main

import (
	"fmt"
	"strings"
)

var text = `Hi there $name.
How is $place? I hope you've been getting a lot of $activity in. Is $someone there? I'm absolutely $expletive going to be there soon.`

func main() {

	fmt.Println(expand(text, func(s string) string {
		// log.Print(s)
		return strings.ToUpper(s[1:])
	}))
}
