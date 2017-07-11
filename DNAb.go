package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	slice, _ := ioutil.ReadFile(`data\rosalind_dna.txt`) //Ans: 14 18 27 20
	fmt.Printf("%s\n", slice)

	for _, r := range "ACGT" {
		sep := []byte(string(r))
		fmt.Printf("%d ", bytes.Count(slice, sep))
	}

	for _, r := range "ACGT" {
		sep := []byte{byte(r)}
		fmt.Printf("%d ", bytes.Count(slice, sep))
	}

	for _, r := range []string{"A", "C", "G", "T"} {
		sep := []byte(r)
		fmt.Printf("%d ", bytes.Count(slice, sep))
	}

}
