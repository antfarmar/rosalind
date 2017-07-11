package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fasta := readFile(`data\rosalind_sseq.txt`)
	dna := getDna(fasta)
	ind := subseq(dna[0], dna[1])
	printsl(ind)
}

func subseq(dna string, key string) []int {
	ind := []int{0}
	for i, c := range key {
		pos := strings.Index(dna, string(c))
		if pos == -1 {
			return nil
		} else {
			pos++
			dna = dna[pos:]
			ind = append(ind, ind[i]+pos)
			continue
		}
	}
	return ind[1:]
}

func printsl(slice []int) {
	println(strings.Trim(fmt.Sprintf("%v", slice), "[]"))
}

func getDna(data []string) []string {
	var buf bytes.Buffer
	for _, line := range data {
		if line == "" || line[0] == '>' {
			buf.WriteString("\n")
		} else {
			buf.WriteString(line)
		}
	}
	return strings.Fields(buf.String())
}

func readFile(file string) []string {
	b, _ := ioutil.ReadFile(file)
	return strings.Fields(string(b))
}
