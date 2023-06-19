package util

import (
	"fmt"
	"math/rand"
)

func sample(cdf []float32) int {
	r := rand.Float32()

	bucket := 0
	for r > cdf[bucket] {
		bucket++
	}
	return bucket
}

func main() {
	// probability density function
	pdf := []float32{0.3, 0.4, 0.2, 0.1}

	// get cdf
	cdf := []float32{0.0, 0.0, 0.0, 0.0}
	cdf[0] = pdf[0]
	for i := 1; i < 4; i++ {
		cdf[i] = cdf[i-1] + pdf[i]
	}

	// test sampling with 100 samples
	samples := []float32{0.0, 0.0, 0.0, 0.0}

	for i := 0; i < 100; i++ {
		samples[sample(cdf)]++
	}

	// normalize
	for i := 0; i < 4; i++ {
		samples[i] /= 100.
	}

	fmt.Println(samples)
	fmt.Println(pdf)

}
