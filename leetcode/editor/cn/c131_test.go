package main

import "testing"

func TestC131(t *testing.T) {
	type testType struct {
		input string
		want  [][]string
	}
	arr := []testType{
		{
			"",
			[][]string{{}},
		},
		{
			"aab",
			[][]string{{"aa", "b"}, {"a", "a", "b"}},
		},
		{
			"abacd",
			[][]string{
				{"a", "b", "c", "d"},
				{"aba", "c", "d"},
			},
		},
		{
			"abacdc",
			[][]string{
				{"a", "b", "a", "c", "d", "c"},
				{"a", "b", "a", "cdc"},
				{"aba", "c", "d", "c"},
				{"aba", "cdc"},
			},
		},
	}
	for key, value := range arr {
		t.Logf("No.%v input:%v, want:%v \n", key, value.input, value.want)
		output := partition(value.input)
		for i := 0; i < len(output); i++ {
			for j := 0; j < len(output[i]); j++ {
				if output[i][j] != value.want[i][j] {
					t.Errorf("Error! ***input %v***want %v***get %v***", value.input, &value.want, &output)
				}
			}
		}
	}
}

//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
