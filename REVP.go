package main

import (
	"bytes"
	"io/ioutil"
	"strings"
)

func main() {
	data := readData(`data\rosalind_revp.txt`)
	dna := parseFasta(data)[1]
	for i := 0; i < len(dna); i += 1 {
		for j := i + 4; j <= i+12 && j <= len(dna); j += 2 {
			if dna[i:j] == rev(cmp(dna[i:j])) {
				println(i+1, j-i)
			}
		}
	}

}

func cmp(dna string) string {
	ATCG := func(r rune) rune {
		switch {
		case r == 'A':
			return 'T'
		case r == 'T':
			return 'A'
		case r == 'C':
			return 'G'
		case r == 'G':
			return 'C'
		}
		return r
	}
	return strings.Map(ATCG, dna)
}

func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
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
