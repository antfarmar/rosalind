package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// Hashtable of DNA {id: [ids of adjacent bases]}
type AdjList map[string][]string

type Dna struct {
	id   string
	base string
}

var (
	re      = regexp.MustCompile(`([ACGT])\s+([ACGT])`) //to join split lines
	adjList = make(AdjList)
)

func main() {
	// Get the data
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_grph.txt`)
	fasta := strings.Fields(re.ReplaceAllString(string(rawBytes), "$1$2"))

	// Compute the answer
	computeAdjList(makeDnaList(fasta), 3)
	fmt.Println(stringifyTable())
}

// Populates hashtable as {id: [ids of adjacent bases]}
func computeAdjList(list []Dna, lng int) {
	for _, dna1 := range list {
		for _, dna2 := range list {
			prefix, suffix := dna1.base[len(dna1.base)-lng:], dna2.base[:lng]
			if prefix == suffix && dna1 != dna2 {
				adjList[dna1.id] = append(adjList[dna1.id], dna2.id)
			}
		}
	}
}

// List of DNA ids & bases
func makeDnaList(fasta []string) (list []Dna) {
	for i := 0; i < len(fasta); i += 2 {
		record := Dna{id: fasta[i][1:], base: fasta[i+1]}
		list = append(list, record)
	}
	return list
}

// Makes the AdjList table printable
func stringifyTable() (out string) {
	for key, val := range adjList {
		for _, v := range val {
			out += fmt.Sprintf("%v %v\n", key, v)
		}
	}
	return
}
