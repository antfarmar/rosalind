package main

import (
	"bytes"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	perms, num := permsMonte(5)
	println(num)
	println(keys(perms))
}

func permsMonte(n int) (map[string]bool, int) {
	perms := make(map[string]bool)
	f := fac(n)
	for len(perms) < f {
		p := rand.Perm(n)
		s := intToStr(p)
		perms[s] = true
	}
	return perms, f
}

func keys(m map[string]bool) string {
	var b bytes.Buffer
	for k := range m {
		b.WriteString(k + "\n")
	}
	return b.String()
}

func intToStr(i []int) string {
	var b bytes.Buffer
	for _, n := range i {
		b.WriteRune('1' + rune(n))
		b.WriteRune(' ')
	}
	return b.String()
}

func fac(n int) int {
	if n == 1 {
		return n
	}
	return n * fac(n-1)
}
