// Recursion
package main

func main() {
	print(fib(30,5)) // == 5164501096416
}

func fib(month, pairs int) int {
	if month == 1 || month == 2 {
		return 1
	}
	return fib(month-1, pairs) + pairs*fib(month-2, pairs)
}


/* Memoization using MAP
package main

func RabbitsCount(start, months, rate int) (count int) {
    counts := map[int]int{1: start, 2: start + start*rate}
    for i := 3; i < months; i++ {
        counts[i] = counts[i-2]*rate + counts[i-1]
    }

    return counts[months-1]
}

func main() {
    println(RabbitsCount(1, 32, 4))
}
*/