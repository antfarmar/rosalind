package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	data := readData(`data\rosalind_subs.txt`)
	dna := data[0]
	tgt := data[1]
	sum := 0
	for {
		pos := strings.Index(dna, tgt)
		if pos == -1 {
			break
		} else {
			sum += pos + 1
			dna = dna[pos+1:]
			print(sum, " ")
		}
	}
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
