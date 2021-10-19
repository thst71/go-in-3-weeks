package main

import (
	"fmt"
	"testing"
)

func TestProverb_String(t *testing.T) {
	pv := proverb{line: "some line"}

	if "text >some line<" != fmt.Sprint(pv) {
		t.Errorf("it is not matching the result")
	}
}

//func Test_charcount(t *testing.T) {
//	type args struct {
//		line string
//	}
//	tests := []struct {
//		name string
//		args args
//		want map[rune]int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := charcount(tt.args.line); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("charcount() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_loadProverbs(t *testing.T) {
//	type args struct {
//		filename string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []*proverb
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := loadProverbs(tt.args.filename)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("loadProverbs() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("loadProverbs() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_newProverb(t *testing.T) {
//	type args struct {
//		line string
//	}
//	tests := []struct {
//		name string
//		args args
//		want *proverb
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := newProverb(tt.args.line); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("newProverb() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_proverb_charStats(t *testing.T) {
//	type fields struct {
//		line  string
//		chars map[rune]int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pv := proverb{
//				line:  tt.fields.line,
//				chars: tt.fields.chars,
//			}
//			if got := pv.charStats(); got != tt.want {
//				t.Errorf("charStats() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_proverb_countChars(t *testing.T) {
//	type fields struct {
//		line  string
//		chars map[rune]int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pv := &proverb{
//				line:  tt.fields.line,
//				chars: tt.fields.chars,
//			}
//		})
//	}
//}
