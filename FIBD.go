package main

func main() {
	println(mortalRabbits(1, 84, 1, 19))
}

func mortalRabbits(start, months, spawn, death int) uint64 {
	counts := map[int]uint64{0: 1, 1: 1, 2: 1}
	for n := 3; n <= months; n++ {
		counts[n] = counts[n-1] + counts[n-2] //*spawn
		if n > death {
			counts[n] -= counts[n-death-1]
		}
	}
	return counts[months]
}
