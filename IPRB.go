package main

import "fmt"

//                           dd dh dr  hh  hr hd  rr rd rh
var punnettProbs = []float64{1, 1, 1, .75, .5, 1, 0, 1, .5}
var dhr = []float64{30, 15, 15}

func main() {
	pairProbs := make([]float64, 9)
	tot := sum(dhr)

	for i := 0; i < len(pairProbs); i += 3 {
		pairProbs[i+0] = dhr[i/3] / tot * (dhr[i/3] - 1) / (tot - 1)
		pairProbs[i+1] = dhr[i/3] / tot * dhr[(i/3+1)%3] / (tot - 1)
		pairProbs[i+2] = dhr[i/3] / tot * dhr[(i/3+2)%3] / (tot - 1)
	}

	for i, pp := range punnettProbs {
		pairProbs[i] *= pp
	}

	fmt.Printf("%v", sum(pairProbs))
}

func sum(slice []float64) (sum float64) {
	for _, f := range slice {
		sum += f
	}
	return
}
