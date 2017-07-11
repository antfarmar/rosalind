package main

import (
	"bytes"
	"io/ioutil"
	"strings"
)

var codon = map[string]string{
	"UUU": "F", "CUU": "L", "AUU": "I", "GUU": "V",
	"UUC": "F", "CUC": "L", "AUC": "I", "GUC": "V",
	"UUA": "L", "CUA": "L", "AUA": "I", "GUA": "V",
	"UUG": "L", "CUG": "L", "AUG": "M", "GUG": "V",
	"UCU": "S", "CCU": "P", "ACU": "T", "GCU": "A",
	"UCC": "S", "CCC": "P", "ACC": "T", "GCC": "A",
	"UCA": "S", "CCA": "P", "ACA": "T", "GCA": "A",
	"UCG": "S", "CCG": "P", "ACG": "T", "GCG": "A",
	"UAU": "Y", "CAU": "H", "AAU": "N", "GAU": "D",
	"UAC": "Y", "CAC": "H", "AAC": "N", "GAC": "D",
	"UAA": " ", "CAA": "Q", "AAA": "K", "GAA": "E",
	"UAG": " ", "CAG": "Q", "AAG": "K", "GAG": "E",
	"UGU": "C", "CGU": "R", "AGU": "S", "GGU": "G",
	"UGC": "C", "CGC": "R", "AGC": "S", "GGC": "G",
	"UGA": " ", "CGA": "R", "AGA": "R", "GGA": "G",
	"UGG": "W", "CGG": "R", "AGG": "R", "GGG": "G",
}

func main() {
	rna := readData(`data\rosalind_prot.txt`)[0]
	println(translate(rna))
}

func translate(rna string) string {
	var b bytes.Buffer
	for i := 0; i <= len(rna)-3; i += 3 {
		b.WriteString(codon[rna[i:i+3]])
	}
	return b.String()
}

func readData(fi string) []string {
	b, _ := ioutil.ReadFile(fi)
	return strings.Fields(string(b))
}
