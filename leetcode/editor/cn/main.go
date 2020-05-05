package main

import "fmt"

func main() {
	fmt.Println("test")
	//longestPalindrome("ababd")
	//fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	//fmt.Println(findMedianSortedArrays([]int{1,2,3}, []int{3,4,5}))
	//fmt.Println(reverse(123))
	//fmt.Println(1 << 31)
	//fmt.Println(-1 << 31)
	s := "1234"
	fmt.Println(s[0:1])
	type testType struct {
		input string
		want  [][]string
	}
	arr := []testType{
		{
			"aab",
			[][]string{{"aa", "b"}, {"a", "a", "b"}},
		},
	}
	output := partition(arr[0].input)
	fmt.Println(&output)
	fmt.Println(&arr[0].want)
}
