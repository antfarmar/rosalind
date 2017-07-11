package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data := readData(`data\rosalind_gc.txt`)
	data = parseFasta(data)
	id, gc := "", 0.0
	for i := 0; i < len(data); i += 2 {
		tmp := gcContent(data[i+1])
		if tmp > gc {
			gc = tmp
			id = data[i][1:]
		}
	}
	fmt.Printf("%s\n%v\n", id, gc*100)
}

func gcContent(dna string) float64 {
	c := strings.Count(dna, "C")
	g := strings.Count(dna, "G")
	return float64(c+g) / float64(len(dna))
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
