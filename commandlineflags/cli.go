package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileNameArg := flag.String("f", "", "use this to specify the input file or set the FILE variable")
	flag.Parse()

	var filename string

	if *fileNameArg == "" {
		filename = os.Getenv("FILE")
		if filename == "" {
			fmt.Println("Neither FILE variable not -f flag are given - exiting")
			os.Exit(-1)
		}
	} else {
		filename = *fileNameArg
	}

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
