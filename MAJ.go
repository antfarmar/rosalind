package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data := readData(`data\rosalind_maj.txt`)
	ints := parseInts(data)
	k, n := ints[0], ints[1]

	for i := 2; i < k*n; i += n {
		fmt.Printf("%d ", majhash(&ints, i, i+n))
	}

	fmt.Print("\n")

	for i := 2; i < k*n; i += n {
		fmt.Printf("%d ", majvote(&ints, i, i+n))
	}
}

// Using Moore's Voting: O(n) time, O(1) space
func majvote(slice *[]int, lo, hi int) int {
	candidate, cnt, maj := (*slice)[lo], 0, -1

	for i := lo; i < hi; i++ {
		element := (*slice)[i]
		if candidate == element {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			candidate, cnt = element, 1
		}
	}

	cnt = 0
	for i := lo; i < hi; i++ {
		if candidate == (*slice)[i] {
			cnt++
			if cnt > (hi-lo)/2 {
				maj = candidate
				break
			}
		}
	}
	return maj
}

// Using hash map: O(n) time, O(n) space
func majhash(slice *[]int, lo, hi int) int {
	maj, amt := -1, hi-lo
	counts := make(map[int]int, amt)
	for i := lo; i < hi; i++ {
		counts[(*slice)[i]]++
	}

	for k, v := range counts {
		if v > amt/2 {
			maj = k
			break
		}
	}
	return maj
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
