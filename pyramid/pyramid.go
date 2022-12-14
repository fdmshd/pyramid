package pyramid

import "sync"

const sizeOfChunk = 4

// f - the function that needs to be performed at each level
func BuildPyramid(f func(intens [][]uint8, lvl int), firstLvl [][]uint8, levels, goVal int) {
	next := firstLvl
	for lvl := 0; lvl < levels; lvl++ {
		f(next, lvl)
		enlarged := Enlarge(next)
		next = NextLvl(enlarged, goVal)
	}
}

func Enlarge(intensity [][]uint8) [][]uint8 {
	n := len(intensity)
	m := len(intensity[0])

	newIntens := make([][]uint8, n+2)
	for i := 0; i < n+2; i++ {
		newIntens[i] = make([]uint8, m+2)
	}
	newIntens[0][0] = intensity[0][0]
	newIntens[0][m+1] = intensity[0][m-1]
	newIntens[n+1][0] = intensity[n-1][0]
	newIntens[n+1][m+1] = intensity[n-1][m-1]

	for i := 1; i < m+1; i++ {
		newIntens[0][i] = intensity[0][i-1]
		newIntens[n+1][i] = intensity[n-1][i-1]
	}
	for i := 1; i < n+1; i++ {
		newIntens[i][0] = intensity[i-1][0]
		newIntens[i][m+1] = intensity[i-1][m-1]
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			newIntens[i][j] = intensity[i-1][j-1]
		}
	}
	return newIntens
}

func NextLvl(enlarged [][]uint8, numGoroutines int) [][]uint8 {
	n := len(enlarged) - 2
	m := len(enlarged[0]) - 2
	res := make([][]uint8, n/2)
	for i := 0; i < n/2; i++ {
		res[i] = make([]uint8, m/2)
	}

	totalChunks := n * m / sizeOfChunk
	if totalChunks < numGoroutines {
		numGoroutines = totalChunks
	}
	chunksPerGoroutine := totalChunks / numGoroutines
	chunksPerRow := m / 2
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for goroutine := 0; goroutine < numGoroutines; goroutine++ {
		go func(thread int) {
			cur := thread * chunksPerGoroutine
			if thread == numGoroutines-1 && chunksPerGoroutine*numGoroutines != totalChunks {
				chunksPerGoroutine += totalChunks - chunksPerGoroutine*numGoroutines
			}
			for chunk := cur; chunk < cur+chunksPerGoroutine; chunk++ {
				x := 2 * (chunk / chunksPerRow)
				y := 2 * (chunk % chunksPerRow)
				res[x/2][y/2] = evalChunk(enlarged, x, y)
			}
			wg.Done()
		}(goroutine)
	}
	wg.Wait()
	return res
}

func evalChunk(mat [][]uint8, x, y int) uint8 {
	sum := 0
	for i := x; i < x+sizeOfChunk; i++ {
		for j := y; j < y+sizeOfChunk; j++ {
			sum += int(mat[i][j])
		}
	}
	return uint8(sum / 16)
}
