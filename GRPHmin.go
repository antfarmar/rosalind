package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`([ACGT])\s+([ACGT])`)

func main() {
	b, _ := ioutil.ReadFile(`data\rosalind_grph.txt`)
	f := strings.Fields(re.ReplaceAllString(string(b), "$1$2"))
	for i := 0; i < len(f); i += 2 {
		for j := 0; j < len(f); j += 2 {
			if i != j && f[i+1][len(f[i+1])-3:] == f[j+1][:3] {
				fmt.Println(f[i][1:], f[j][1:])
			}

		}
	}
}
