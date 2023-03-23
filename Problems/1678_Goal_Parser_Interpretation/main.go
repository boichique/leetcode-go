package main

import (
	"fmt"
	"strings"
)

func main() {
	command := "(al)G(al)()()G"
	fmt.Println(interpret(command))
}

func interpret(command string) string {
	r := strings.NewReplacer("()", "o", "(al)", "al")
	return r.Replace(command)
}
