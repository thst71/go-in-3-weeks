package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Printf("%s <filename>", args[0])
		os.Exit(-1)
	}

	filename := args[1]

	txt, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Printf("file not found %s\n", err)
		os.Exit(-2)
	}

	lines := strings.Split(string(txt), "\n")

	for _, line := range lines {
		cmap := charcount(line)
		fmt.Println(line)
		for chr, cnt := range cmap {
			fmt.Printf("'%c' = %d|", chr, cnt)
		}
		fmt.Println()
	}
}

func charcount(line string) map[rune]int {
	cmap := make(map[rune]int, 0)
	for _, char := range line {
		cmap[char] = cmap[char] + 1
	}
	return cmap
}
