package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	data := readData(`data\rosalind_hamm.txt`)
	dna1 := data[0]
	dna2 := data[1]
	dst := 0
	for i := range dna1 {
		if dna1[i] != dna2[i] {
			dst += 1
		}
	}
	print(dst, "\n")
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
