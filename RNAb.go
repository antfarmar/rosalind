// RNAb.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	byt, _ := ioutil.ReadFile(`data\rosalind_rna.txt`)
	fmt.Printf("%s", byt)
	old, new := []byte("T"), []byte("U")
	fmt.Printf("%s", bytes.Replace(byt, old, new, -1))
}
