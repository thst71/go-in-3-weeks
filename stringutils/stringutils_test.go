package stringutils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	var casesText string

	if cases, err := ioutil.ReadFile("testcases.txt"); err != nil {
		log.Fatalln("testcase file not found")
	} else {
		casesText = string(cases)
	}

	for idx, line := range strings.Split(casesText, "\n") {
		fmt.Printf("adding %d to %s\n", idx, line)
	}

	code := m.Run()
	os.Exit(code)
}

func BenchmarkLcase5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lcase("THISI")
	}
}

func BenchmarkLcase10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lcase("THIS IS A ")
	}
}

func BenchmarkLcase20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lcase("THIS IS  COMMON TEXT")
	}
}

func BenchmarkLcase40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lcase("THIS IS  COMMON TEXTTHIS IS  COMMON TEXT")
	}
}

func TestLcase(t *testing.T) {
	type args struct {
		incoming string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"it should convert to lower case", args{"UPPER CASE"}, "upper case"},
		{"it should convert the empty string to lower case", args{""}, ""},
		{"it should convert symstring to lower case", args{"UPPÈR1234!§$%&CASE"}, "uppèr1234!§$%&case"},
		{"it should convert mixed to lower case", args{"Upper1234!§$%&casE"}, "upper1234!§$%&case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcase(tt.args.incoming); got != tt.want {
				t.Errorf("Lcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUcase(t *testing.T) {
	type args struct {
		incoming string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"it should convert to upper case", args{"upper case"}, "UPPER CASE"},
		{"it should convert the empty string to upper case", args{""}, ""},
		{"it should convert symstring to upper case", args{"uppèr1234!§$%&case"}, "UPPÈR1234!§$%&CASE"},
		{"it should convert mixed to upper case", args{"Upper1234!§$%&casE"}, "UPPER1234!§$%&CASE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ucase(tt.args.incoming); got != tt.want {
				t.Errorf("Ucase() = %v, want %v", got, tt.want)
			}
		})
	}
}
