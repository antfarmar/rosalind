package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
) //comment

func main() {
	data := readData(`data\rosalind_cons.txt`)
	dna := parseFasta(data)
	mat := dnaMatrix(dna)
	pro := profile(mat)
	con := consensus(pro)
	fmt.Println(con)
	fmt.Println(proPrint(pro))
}

func proPrint(pro [][]int) string {
	var b bytes.Buffer
	for i, chr := range "ACGT" {
		row := strings.Trim(fmt.Sprintf("%v", pro[i]), "[]")
		b.WriteRune(chr)
		b.WriteRune(':')
		b.WriteString(" " + row + "\n")
	}
	return b.String()
}

func consensus(pro [][]int) string {
	con := ""
	for col := range pro[0] {
		max, loc := 0, 0
		for row := range pro {
			if pro[row][col] > max {
				max = pro[row][col]
				loc = row
			}
		}
		con += string("ACGT"[loc])
	}
	return con
}

func profile(mat [][]rune) [][]int {
	lenDNA := len(mat[0])
	pro := make([][]int, 4)
	for i, chr := range "ACGT" {
		pro[i] = make([]int, lenDNA)
		for col := range pro[i] {
			for row := range mat {
				if mat[row][col] == chr {
					pro[i][col]++
				}
			}
		}
	}
	return pro
}

func dnaMatrix(fasta []string) [][]rune {
	var mat [][]rune
	for i := 1; i < len(fasta); i += 2 {
		mat = append(mat, []rune(fasta[i]))
	}
	return mat
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
