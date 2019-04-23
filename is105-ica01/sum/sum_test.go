package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{1, 2, 3},
	{4, 5, 9},
	{120, 1, 119},
}

var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{32, 2, 34},
	{4, 183, 188},
	{120, 1, 121},
}

var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{-4, 10, 6},
	{4, 8, 12},
	{-100, 120, 5},
}

var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{8, 10, 18},
	{4, 8, 12},
	{-100, 120, 20},
}

var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{8.1, 10.0, 18.1},
	{4.0, 8.0, 12.0},
	{-100.0, 120.0, 20.0},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
		}
	}
}