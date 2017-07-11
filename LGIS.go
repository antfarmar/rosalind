package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//seq := readIntSeq()
	seq := []int{3, 5, 7, 9, 2, 6, 1, 8, 4}
	n := len(seq)

	memo := makeMemo(n, n)
	seqSorted := make([]int, n)

	copy(seqSorted, seq)
	sort.Ints(seqSorted)

	// Longest Increasing Subsequence
	longest := lcs(seq, seqSorted, memo)
	printlcs(seq, longest, n, n)
	fmt.Println()
	showMemo(memo)

	reset(memo)
	reverse(seqSorted)

	// Longest Decreasing Subsequence
	longest = lcs(seq, seqSorted, memo)
	printlcs(seq, longest, n, n)
	fmt.Println()
	showMemo(memo)
}

func showMemo(m [][]int) {
	for _, i := range m {
		fmt.Println(i)
	}
}

func reverse(seq []int) {
	for i, j := 0, len(seq)-1; i < len(seq)/2; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
}

func makeMemo(m, n int) [][]int {
	var memo [][]int = make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	return memo
}

func reset(memo [][]int) {
	for i := range memo {
		for j := range memo {
			memo[i][j] = 0
		}
	}
}

func lcs(s1, s2 []int, memo [][]int) [][]int {
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			switch {
			case s1[i-1] == s2[j-1]: // match
				memo[i][j] = memo[i-1][j-1] + 1
			case memo[i][j-1] >= memo[i-1][j]:
				memo[i][j] = memo[i][j-1] + 0
			default: // memo[i-1][j] > memo[i][j-1]
				memo[i][j] = memo[i-1][j] + 0
			}
		}
	}
	return memo
}

func printlcs(s []int, r [][]int, i, j int) {
	if i == 0 || j == 0 {
		return
	}
	switch {
	case r[i][j] > r[i-1][j] && r[i][j] > r[i][j-1]:
		printlcs(s, r, i-1, j-1)
		fmt.Printf("%v ", s[i-1])
	case r[i][j-1] >= r[i-1][j]:
		printlcs(s, r, i, j-1)
	default:
		printlcs(s, r, i-1, j)
	}
}

func readIntSeq() (nums []int) {
	rawBytes, _ := ioutil.ReadFile(`data\rosalind_lgis.txt`)
	str := strings.Fields(string(rawBytes))[1:]
	for _, n := range str {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}
	return nums
}
