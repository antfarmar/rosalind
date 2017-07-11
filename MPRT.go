package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var motif = regexp.MustCompile("N[^P][ST][^P]")

func main() {
	idNames := readFile(`data\rosalind_mprt.txt`)
	for _, id := range idNames {
		url := "http://www.uniprot.org/uniprot/" + id + ".fasta"
		fasta := readHttp(url)
		seq := getSeq(fasta)[0]
		loc := motif_locs(seq)
		if loc != nil {
			println(id)
			println(strings.Trim(fmt.Sprintf("%v", loc), "[]"))
		}
	}
}

func motif_locs(seq string) []int {
	//idx := motif.FindAllStringIndex(seq, -1) // Fails on overlaps
	var locs []int
	pos := 0
	idx := motif.FindStringIndex(seq)
	for ; idx != nil; idx = motif.FindStringIndex(seq[pos:]) {
		pos += idx[0] + 1
		locs = append(locs, pos)
	}
	return locs
}

func getSeq(data []string) []string {
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
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}

func readHttp(url string) []string {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(contents), "\n")
}
