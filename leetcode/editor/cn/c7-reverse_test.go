package main

import "testing"

func TestReverse(t *testing.T) {
	type testType struct {
		input int
		want  int
	}
	arr := []testType{
		{123, 321},
		{345, 543},
		{0, 0},
		//1 << 31 = 2^31 = 2147483648
		{1<<31 - 1, 0},
		{-1 << 31, 0},
	}
	for key, value := range arr {
		t.Logf("No.%v input:%v, want:%v \n", key, value.input, value.want)
		output := reverse(value.input)
		if output != value.want {
			t.Errorf("Error! ***input %v***want %v***get %v***", value.input, value.want, output)
		}
	}
}
