package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	slice := readData(`data\rosalind_dna.txt`) //Ans: 14 18 27 20
	for _, s := range slice {
		for _, c := range "ACGT" {
			print(strings.Count(s, string(c)), " ")
		}
	}
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
