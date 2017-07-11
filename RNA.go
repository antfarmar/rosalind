package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	str := readData(`data\rosalind_rna.txt`)[0]
	print(strings.Replace(str, "T", "U", -1))
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
