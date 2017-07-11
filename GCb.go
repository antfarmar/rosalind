package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	data := readBytes(`data\rosalind_gc.txt`)
	data = parseFasta(data)
	//fmt.Printf("%s", data)
	id, gc := []byte{}, 0.0
	for i := 0; i < len(data); i += 2 {
		tmp := gcContent(data[i+1])
		if tmp > gc {
			gc = tmp
			id = data[i][1:]
		}
	}
	fmt.Printf("%s\n%v\n", id, gc*100)
}

func gcContent(dna []byte) float64 {
	c := bytes.Count(dna, []byte("C"))
	g := bytes.Count(dna, []byte("G"))
	return float64(c+g) / float64(len(dna))
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
