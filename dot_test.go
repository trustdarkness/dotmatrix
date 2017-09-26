package main

import (
	"testing"
	"reflect"
)

func Test2x2dot2x2(t *testing.T) {
	m1 := matrix{2, 2, true, []int{1, 2, 3, 4}}
	m2 := matrix{2, 2, true, []int{3, 4, 5, 6}}
	good := matrix{2, 2, true, []int{13,16,29,36}}
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("2x2dot2x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}

func Test4x2dot2x4(t *testing.T) {
	m1 := matrix{2, 4, true, []int{3, 4, 5, 6, 7, 8, 9, 10}}
	m2 := matrix{4, 2, true, []int{1, 2, 3, 4, 5, 6, 7, 8}}
	good := matrix{2, 2, true, []int{82, 100, 146, 180}}
	product := Product(m1, m2)
	if !reflect.DeepEqual(product, good) {
		t.Error("4x2dot2x2 failed")
		t.Error("Expected:")
		PrintMatrix(good)
		t.Error("Got:")
		PrintMatrix(product)
	}
}