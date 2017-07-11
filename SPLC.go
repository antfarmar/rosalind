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
	data := readData(`data\rosalind_splc.txt`)
	dna := parseFasta(data)
	exo := exons(dna[0], dna[1:])
	rna := rna(exo)
	println(translate(rna))
}

func rna(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func exons(dna string, introns []string) string {
	exo := dna
	for _, intr := range introns {
		exo = strings.Replace(exo, intr, "", -1)
	}
	return exo
}

func translate(rna string) string {
	var b bytes.Buffer
	for i := 0; i <= len(rna)-3; i += 3 {
		b.WriteString(codon[rna[i:i+3]])
	}
	return b.String()
}
func parseFasta(data []string) []string {
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
