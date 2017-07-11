package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := readData(`data\rosalind_ins.txt`)
	nums := parseInts(data[1:])
	fmt.Println(nums)
	nums, swaps := insSort(nums)
	fmt.Println(nums)
	fmt.Println(swaps)
}

func insSort(a []int) ([]int, int) {
	swaps := 0
	for i := 1; i < len(a); i++ {
		for k := i; k > 0 && a[k] < a[k-1]; k-- {
			a[k], a[k-1] = a[k-1], a[k]
			swaps++
		}
	}
	return a, swaps
}

func parseInts(ss []string) (nums []int) {
	for _, v := range ss {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	return
}

func readData(name string) []string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return strings.Fields(string(b))
}
