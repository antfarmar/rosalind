package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	byt, _ := ioutil.ReadFile(`data\rosalind_revc.txt`)
	fmt.Printf("%s", byt)
	fmt.Printf("%s\n", comp(rev(byt)))
}

func comp(b []byte) []byte {
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
	return bytes.Map(ATCG, b)
}

func rev(r []byte) []byte {
	//r := bytes.Runes(b)
	i, j := 0, len(r)-1
	for i < len(r)/2 {
		r[i], r[j] = r[j], r[i]
		i, j = i+1, j-1
	}
	return r
}
