package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type ByLength []string

func (a ByLength) Len() int           { return len(a) }
func (a ByLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func main() {
	data := readData(`data\rosalind_lcsm.txt`)
	dna := getDNA(data)
	sort.Sort(ByLength(dna))
	fmt.Println(lcs(dna))
}

func lcs(dna []string) string {
	key := dna[0]
	long := ""
	for size := 1; size <= len(key); size++ {
		for start := 0; start <= len(key)-size; start++ {
			sub := key[start : start+size]
			if hasSub(dna[1:], sub) {
				long = sub
				break
			}
		}
	}
	return long
}

func hasSub(slice []string, sub string) bool {
	for _, str := range slice {
		if !strings.Contains(str, sub) {
			return false
		}
	}
	return true
}

func getDNA(data []string) []string {
	var buf bytes.Buffer
	for _, line := range data {
		if line[0] == '>' {
			buf.WriteString("\n")
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
