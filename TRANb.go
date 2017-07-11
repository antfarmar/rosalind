package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	data := readBytes(`data\rosalind_tran.txt`)
	dna := parseFasta(data)
	fmt.Printf("%v\n", tratio(dna[1], dna[3])) // 1.7065217391304348
}

func tratio(s, t []byte) float64 {
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

func parseFasta(data [][]byte) [][]byte {
	var buf bytes.Buffer
	for _, line := range data {
		if line[0] == '>' {
			buf.WriteRune('\n')
			buf.Write(line)
			buf.WriteRune('\n')
		} else {
			buf.Write(line)
		}
	}
	return bytes.Fields(buf.Bytes())
}

func readBytes(name string) [][]byte {
	b, _ := ioutil.ReadFile(name)
	return bytes.Fields(b)
}
