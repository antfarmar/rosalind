package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	str := readData(`data\rosalind_revc.txt`)[0]
	ATCG := func(r rune) rune {
		switch {
		case r == 'A':
			return 'T'
		case r == 'T':
			return 'A'
		case r == 'C':
			return 'G'
		case r == 'G':
			return 'C'
		}
		return r
	}
	fmt.Println(str)
	fmt.Println(strings.Map(ATCG, rev(str)))
}

func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
