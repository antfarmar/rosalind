package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`([ACGT])\s+([ACGT])`)

func main() {
	// Get the data
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_grph.txt`)
	fasta := strings.Fields(re.ReplaceAllString(string(rawBytes), "$1$2"))
	fmt.Printf("%q\n", fasta)
}
