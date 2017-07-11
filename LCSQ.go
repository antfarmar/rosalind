package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fasta := readFile(`data\rosalind_lcsq.txt`)
	dna := getDna(fasta)
	s, t := dna[0], dna[1]
	m, n := len(dna[0]), len(dna[1])
	lcs := lcs(s, t)
	printlcs(s, lcs, m, n)
}

func lcs(s1, s2 string) [][]int {
	m, n := len(s1), len(s2)
	memo := make([][]int, m+1)

	for i, _ := range memo {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case s1[i-1] == s2[j-1]:
				memo[i][j] = memo[i-1][j-1] + 1
			case memo[i][j-1] >= memo[i-1][j]:
				memo[i][j] = memo[i][j-1]
			default:
				memo[i][j] = memo[i-1][j]
			}
		}
	}
	return memo
}

func printlcs(s string, r [][]int, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	switch {
	case r[i][j] > r[i-1][j] && r[i][j] > r[i][j-1]:
		printlcs(s, r, i-1, j-1)
		fmt.Printf("%c", s[i-1])
	case r[i][j-1] >= r[i-1][j]:
		printlcs(s, r, i, j-1)
	default:
		printlcs(s, r, i-1, j)
	}
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
