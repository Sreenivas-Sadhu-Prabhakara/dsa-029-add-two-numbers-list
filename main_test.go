package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		a, b, want []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}}, // 342+465=807
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9}, []int{1}, []int{0, 0, 0, 0, 1}}, // carry ripples past the end
		{[]int{9, 9}, []int{9, 9, 9}, []int{8, 9, 0, 1}},    // 99+999=1098, unequal lengths
		{[]int{5}, []int{5}, []int{0, 1}},
	}
	for _, c := range cases {
		got := toSlice(AddTwoNumbers(fromSlice(c.a), fromSlice(c.b)))
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("AddTwoNumbers(%v,%v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
