package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type proverb struct {
	line  string
	chars map[rune]int
}

func (pv proverb) String() string {
	return fmt.Sprintf("text >%s<", pv.line)
}

func (pv *proverb) countChars() {
	pv.chars = charcount(pv.line)
}

func newProverb(line string) *proverb {
	newPv := proverb{line: line}
	newPv.countChars()
	return &newPv
}

func charcount(line string) map[rune]int {
	cmap := make(map[rune]int, 0)
	for _, char := range line {
		cmap[char] = cmap[char] + 1
	}
	return cmap
}

func loadProverbs(filename string) ([]*proverb, error) {
	if txt, err := ioutil.ReadFile(filename); err != nil {
		fmt.Printf("file not found %s\n", err)
		return nil, err
	} else {
		lines := strings.Split(string(txt), "\n")
		proverbs := make([]*proverb, 0)
		for _, line := range lines {
			proverbs = append(proverbs, newProverb(line))
		}

		return proverbs, nil
	}
}

func (pv proverb) charStats() string {
	var out string = "|"
	for chr, cnt := range pv.chars {
		out = fmt.Sprintf("%s'%c' = %d|", out, chr, cnt)
	}

	return out
}

func main() {
	fileNameArg := flag.String("f", "", "use this to specify the input file or set the FILE variable")
	flag.Parse()

	var filename string

	if *fileNameArg == "" {
		filename = os.Getenv("FILE")
		if filename == "" {
			fmt.Println("Neither FILE variable nor -f flag is given - exiting")
			os.Exit(-1)
		}
	} else {
		filename = *fileNameArg
	}

	if pvs, err := loadProverbs(filename); err == nil {
		for _, pv := range pvs {
			fmt.Printf("%s\n%s\n", pv.line, pv.charStats())
		}
	}
}
