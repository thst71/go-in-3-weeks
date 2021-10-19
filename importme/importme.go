package main

import (
	"fmt"
	"github.com/thst71/parseplan/go/parser"
)

func main() {
	var p parser.Parser

	p.Greeting = "hurra!"

	fmt.Println(p)
}
