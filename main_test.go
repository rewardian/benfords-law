package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	map1 = map[int]int{
		1: 5000,
		8: 45,
		7: 15,
		9: 600,
		4: 1200,
		0: 700,
	}
)

func TestBenfordValidator(t *testing.T) {
	values := []struct {
		percent float64
		want    bool
	}{
		{30.00, true},
		{35.00, true},
		{25.00, false},
	}

	for _, test := range values {
		testname := fmt.Sprintf("%f", test.percent)

		t.Run(testname, func(t *testing.T) {
			answer := benfordValidator(test.percent)
			if answer != test.want {
				t.Errorf("Got %t, wanted %t", answer, test.want)
			}
		})
	}
}

func TestRemoveIndex(t *testing.T) {
	values := []struct {
		slice []int
		index int
		want  []int
	}{
		{[]int{0, 1, 2, 3}, 1, []int{0, 2, 3}},
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 5, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range values {
		testname := fmt.Sprintf("%v", test.slice)

		t.Run(testname, func(t *testing.T) {
			answer := removeIndex(test.slice, test.index)

			if !reflect.DeepEqual(answer, test.want) {
				t.Errorf("Got %v, wanted %v", answer, test.want)
			}
		})
	}
}

func TestRetrieveFirstDigit(t *testing.T) {
	values := []struct {
		record string
		want   int
	}{
		{"0.051", 5},
		{"-0.051", 5},
		{"-17", 1},
		{".00310", 3},
		{"Bobby Tables", 0},
		{".0098", 9},
	}

	for _, test := range values {
		testname := fmt.Sprintf("%v", test.record)

		t.Run(testname, func(t *testing.T) {
			answer := retrieveFirstDigit(test.record)

			if answer != test.want {
				t.Errorf("Got %d, wanted %d", answer, test.want)
			}
		})
	}
}

func TestSanitizeColumnValue(t *testing.T) {
	values := []struct {
		column string
		want   int
	}{
		{"7", 6},
		{"1", 0},
		{"4", 3},
	}

	for _, test := range values {
		testname := fmt.Sprintf("%v", test.column)

		t.Run(testname, func(t *testing.T) {
			answer := sanitizeColumnValue(test.column)

			if answer != test.want {
				t.Errorf("Got %d, wanted %d", answer, test.want)
			}
		})
	}

}

func TestSortMap(t *testing.T) {
	want := []int{0, 0, 0, 0, 1, 4, 7, 8, 9}

	testname := fmt.Sprintf("%s", "map1")

	t.Run(testname, func(t *testing.T) {
		answer := sortMap(map1)
		if !reflect.DeepEqual(answer, want) {
			t.Errorf("Got %v, wanted %v", answer, want)
		}
	})

}
