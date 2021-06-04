package karateChop_test

import (
	karateChop "github.com/bshyn/go-katas/day_01"
	"testing"
)

func TestBinaryChop(t *testing.T) {
	var emptySlice []int
	emptyResult := karateChop.BinaryChop(1, emptySlice)
	if emptyResult != -1 {
		t.Errorf("BinaryChop with empty slice; Expected -1 got %d", emptyResult)
	}

	singleElementSlice := []int{1}
	singleElementResult := karateChop.BinaryChop(1, singleElementSlice)
	if singleElementResult != 0 {
		t.Errorf("BinaryChop with single element slice; Expected 0 got %d", singleElementResult)
	}

	twoElementsSlice := []int{3,25}
	twoElementsResult := karateChop.BinaryChop(25, twoElementsSlice)
	if twoElementsResult != 1 {
		t.Errorf("BinaryChop with two elements slice; Expected 1 got %d", twoElementsResult)
	}

	threeElementsSlice := []int{9, 11, 25}
	threeElementsResult := karateChop.BinaryChop(9, threeElementsSlice)
	if threeElementsResult != 0 {
		t.Errorf("BinaryChop with three elements slice to the left; Expected 0 got %d", threeElementsResult)
	}

	threeElementsResult = karateChop.BinaryChop(25, threeElementsSlice)
	if threeElementsResult != 2 {
		t.Errorf("BinaryChop with three elements slice to the right; Expected 2 got %d", threeElementsResult)
	}

	fourElementsSlice := []int{-5, 0, 3, 7}
	fourElementsResult := karateChop.BinaryChop(-5, fourElementsSlice)
	if fourElementsResult != 0 {
		t.Errorf("BinaryChop with four elements slice to the left; Expected 0 got %d", fourElementsResult)
	}

	fourElementsResult = karateChop.BinaryChop(7, fourElementsSlice)
	if fourElementsResult != 3 {
		t.Errorf("BinaryChop with four elements slice to the right; Expected 3 got %d", fourElementsResult)
	}

	lotsOfElementsSlice := []int {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	lotsOfElementsResult := karateChop.BinaryChop(9, lotsOfElementsSlice)
	if lotsOfElementsResult != 9 {
		t.Errorf("BinaryChop with lots of elements; Expected 9 got %d", lotsOfElementsResult)
	}
}

func BenchmarkBinaryChop(b *testing.B) {
	slice := []int {0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	for i := 0; i < b.N; i++ {
		_ = karateChop.BinaryChop(0, slice)
	}
}
