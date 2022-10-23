package pyramid

import (
	"reflect"
	"testing"
)

func TestEnlarge(t *testing.T) {
	matrix := [][]uint8{{55, 12}, {44, 196}}
	want := [][]uint8{{55, 55, 12, 12}, {55, 55, 12, 12}, {44, 44, 196, 196}, {44, 44, 196, 196}}
	got := Enlarge(matrix)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Enlarge(%v) = %v, want:%v", matrix, got, want)
	}
}

func TestNextLvl(t *testing.T) {
	enarged := [][]uint8{
		{4, 4, 4, 8, 6, 6},
		{4, 4, 4, 8, 6, 6},
		{5, 5, 6, 8, 8, 8},
		{6, 6, 7, 10, 11, 11},
		{7, 7, 8, 12, 13, 13},
		{7, 7, 8, 12, 13, 13},
	}
	want := [][]uint8{
		{5, 7},
		{7, 10},
	}
	for i := 1; i < 16; i++ {
		got := NextLvl(enarged, i)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("NextLlv = %v, want:%v", got, want)
		}
	}
}

func TestEvalChunk(t *testing.T) {
	chunk := [][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	got := evalChunk(chunk, 0, 0)
	want := uint8(7)
	if got != want {
		t.Errorf("EvalChunk = %v, want:%v", got, want)
	}

}
