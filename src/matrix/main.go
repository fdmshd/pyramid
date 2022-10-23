package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"pyramid/pyramid"
	"time"
)

var dFlag = flag.Int("D", 12, "number of rows of the matrix = 2^D")
var sFlag = flag.Int("S", 12, "number of columns of the matrix = 2^D")
var nFlag = flag.Int("N", 6, "number of pyramid levels")

func main() {
	flag.Parse()
	d := int(math.Pow(2, float64(*dFlag)))
	s := int(math.Pow(2, float64(*sFlag)))
	intensity := make([][]uint8, d)
	for i := 0; i < d; i++ {
		intensity[i] = make([]uint8, s)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < d; i++ {
		for j := 0; j < s; j++ {
			intensity[i][j] = uint8(rand.Intn(256))
		}
	}

	for threadN := 1; threadN <= 16; threadN++ {
		var timeAv time.Duration
		for i := 0; i < 3; i++ {
			next := intensity
			start := time.Now()
			for lvl := 0; lvl < *nFlag; lvl++ {
				enlarged := pyramid.Enlarge(next)
				next = pyramid.NextLvl(enlarged, threadN)
			}
			timeAv += time.Since(start)
		}
		timeAv = timeAv / 3
		fmt.Printf("goroutines:%d - average time:%v \n", threadN, timeAv)
	}

}
