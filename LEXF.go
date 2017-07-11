package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_lexf.txt`)
	data := strings.Fields(string(rawBytes))

	symbols := data[:len(data)-1]
	n := int(data[len(data)-1][0] - '0')

	for _, kmer := range cartProd(symbols, n) {
		fmt.Println(kmer)
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
