package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`([ACGT])\s+([ACGT])`)

func main() {
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_kmer.txt`)
	dna := strings.Fields(re.ReplaceAllString(string(rawBytes), "$1$2"))[1]

	for _, kmer := range cartProd([]string{"A", "C", "G", "T"}, 4) {
		cnt, pos := 0, 0
		for index := strings.Index(dna, kmer); index >= 0; cnt++ {
			pos += index + 1
			index = strings.Index(dna[pos:], kmer)
		}
		fmt.Printf("%v ", cnt)
	}
}

func cartProd(abc []string, n int) (nary []string) {
	nary = abc
	for ; n > 1; n-- {
		temp := []string{}
		for _, s := range abc {
			for _, t := range nary {
				temp = append(temp, s+t)
			}
		}
		nary = temp
	}
	return
}
