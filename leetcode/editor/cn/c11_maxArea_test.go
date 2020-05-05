package main

import "testing"

func TestMaxArea(t *testing.T) {

	type testType struct {
		input []int
		want  int
	}
	arr := []testType{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 3}, 2},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{3, 1, 2, 4}, 9},
		{[]int{1, 1, 3, 3, 4, 4}, 9},
	}

	for key, value := range arr {
		t.Logf("No.%v input:%v, want:%v \n", key, value.input, value.want)
		output := maxArea(value.input)
		if output != value.want {
			t.Errorf("Error! ***input %v***want %v***get %v***", value.input, value.want, output)
		}
	}
}
