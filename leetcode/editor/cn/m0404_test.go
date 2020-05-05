package main

import "testing"

func TestRomanToInt(t *testing.T) {
	type testType struct {
		input *TreeNode
		want  bool
	}
	arr := []testType{
		{test, 3},
		{"IV", 4},
		{"IX", 9},
		//1 << 31 = 2^31 = 2147483648
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for key, value := range arr {
		t.Logf("No.%v input:%v, want:%v \n", key, value.input, value.want)
		output := romanToInt(value.input)
		if output != value.want {
			t.Errorf("Error! ***input %v***want %v***get %v***", value.input, value.want, output)
		}
	}
}

//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
