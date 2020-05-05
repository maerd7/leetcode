package main

import "testing"

func TestIntToRoman(t *testing.T) {
	type testType struct {
		input int
		want  string
	}
	arr := []testType{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		//1 << 31 = 2^31 = 2147483648
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for key, value := range arr {
		t.Logf("No.%v input:%v, want:%v \n", key, value.input, value.want)
		output := intToRoman(value.input)
		if output != value.want {
			t.Errorf("Error! ***input %v***want %v***get %v***", value.input, value.want, output)
		}
	}
}

//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
