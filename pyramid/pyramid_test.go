package pyramid

import (
	"testing"
)

func compareMatrices(a, b [][]uint8) bool {
	n := len(a)
	m := len(a[0])
	if n != len(b) || m != len(b[0]) {
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestEnlarge(t *testing.T) {
	matrix := [][]uint8{{55, 12}, {44, 196}}
	want := [][]uint8{{55, 55, 12, 12}, {55, 55, 12, 12}, {44, 44, 196, 196}, {44, 44, 196, 196}}
	got := Enlarge(matrix)

	if !compareMatrices(got, want) {
		t.Errorf("Enlarge(%v) = %v, want:%v", matrix, want, got)
	}
}

func TestNextLvl(t *testing.T) {

}
