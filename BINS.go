package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := readData(`data\rosalind_bins.txt`)
	ints := parseInts(data)
	n := ints[0]
	nums := ints[2 : 2+n]
	keys := ints[2+n:]
	fmt.Println(nums, keys)

	locs := make([]int, len(keys))
	for j, k := range keys {
		locs[j] = binSearch(&nums, 0, n, k)
	}
	fmt.Println(locs)
}

func binSearch(nums *[]int, lo, hi, key int) int {
	if lo == hi {
		return -1
	}
	mid := (lo + hi) / 2
	val := (*nums)[mid]
	if key < val {
		return binSearch(nums, lo, mid, key)
	} else if key > val {
		return binSearch(nums, mid+1, hi, key)
	} else { // key==val
		return mid + 1
	}
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
