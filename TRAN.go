package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data := readData(`data\rosalind_tran.txt`)
	dna := parseFasta(data)
	fmt.Printf("%v\n", tratio(dna[1], dna[3])) // 1.7065217391304348
}

func tratio(s, t string) float64 {
	transitions := 0.0
	transversions := 0.0
	for i := range s {
		if s[i] != t[i] {
			if s[i]+t[i] == 'A'+'G' || s[i]+t[i] == 'C'+'T' {
				transitions++
			} else {
				transversions++
			}
		}
	}
	return transitions / transversions
}

func parseFasta(data []string) []string {
	var buf bytes.Buffer
	for _, line := range data {
		if line[0] == '>' {
			buf.WriteString("\n" + line + "\n")
		} else {
			buf.WriteString(line)
		}
	}
	return strings.Fields(buf.String())
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
