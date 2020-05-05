package main

import "fmt"

func main() {
	s := "1234"
	fmt.Println(len(s))
	//checkA := [len(s)][len(s)]int{}
	checkA := make([][]int, len(s))
	for i := 0; i < len(checkA); i++ {
		checkA[i] = make([]int, len(s))
	}
	fmt.Println(checkA[0][0])
	arr := [][]int{{1, 2, 3}, {3, 4, 5}}
	//add(&arr)
	mul(arr)
	fmt.Println(arr)

}

func add(p *[][]int) {
	(*p)[1][1] = 10
}

func mul(b [][]int) {
	b[0][0] = 10
}
