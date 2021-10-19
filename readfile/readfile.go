package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	txt, err := ioutil.ReadFile("readfile/verbs.txt")

	if err != nil {
		fmt.Printf("file not found %s\n", err)
		return
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
