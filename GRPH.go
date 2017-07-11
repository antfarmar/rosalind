package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type Dna struct {
	id   string
	code string
}

func main() {
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_grph.txt`)
	fasta := parse(strings.Fields(string(rawBytes)))
	//fmt.Println(fasta)
	idDnaList := makeList(fasta)
	//fmt.Println(idDnaList)
	adjList_O3 := adjList(idDnaList, 3)
	fmt.Println(stringifyTable(adjList_O3))
}

func adjList(list []Dna, lng int) (table map[string][]string) {
	table = make(map[string][]string)
	for _, dna1 := range list {
		suffix := dna1.code[len(dna1.code)-lng:]
		for _, dna2 := range list {
			if dna1 == dna2 {
				continue
			}
			prefix := dna2.code[:lng]
			if suffix == prefix {
				table[dna1.id] = append(table[dna1.id], dna2.id)
			}
		}
	}
	return
}

func makeList(fasta []string) (list []Dna) {
	for i := 0; i < len(fasta); i += 2 {
		record := Dna{id: fasta[i][1:], code: fasta[i+1]}
		list = append(list, record)
	}
	return list
}

func stringifyTable(table map[string][]string) (out string) {
	for key, val := range table {
		for _, v := range val {
			out += fmt.Sprintf("%v %v\n", key, v)
		}
	}
	return
}

func parse(fasta []string) []string {
	var buf bytes.Buffer
	for _, line := range fasta {
		if line[0] == '>' || line == "" {
			buf.WriteString("\n" + line + "\n")
		} else {
			buf.WriteString(line)
		}
	}
	return strings.Fields(buf.String())
}
