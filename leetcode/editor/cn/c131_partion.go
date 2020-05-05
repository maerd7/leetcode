package main

import "fmt"

func find(s string, tree *Tree) {

	if len(s) == 0 {
		return
	} else if len(s) == 1 {
		tempTree := &Tree{
			id: s[0:1],
		}
		tree.children = append(tree.children, tempTree)
		return
	}
	//find all reverse
	for i := 0; i < len(s); i++ {
		//依次判断是不是回文
		isReverse := true
		for j := 0; j <= i; j++ {
			if s[j] != s[i-j] {
				isReverse = false
				break
			}
		}
		// 是回文, 则依次递归调用
		if isReverse {
			tempTree := &Tree{
				id: s[0 : i+1],
			}
			// 接到入参树之后
			tree.children = append(tree.children, tempTree)
			// i = len(s) - 1是最后一位，等于到时候就不可以s[i+1:]了
			if i+1 < len(s) {
				find(s[i+1:], tempTree)
			}
		}
	}
}

type Tree struct {
	id       string
	children []*Tree
}

func dfs(tree *Tree) {
	if tree.children == nil {
		fmt.Println("-")
	}
	if tree.id != "root" {
		fmt.Print(tree.id)
	}
	for i := 0; i < len(tree.children); i++ {
		dfs(tree.children[i])
	}
}

// 5.4 22:48
func partition(s string) [][]string {
	tree := &Tree{
		id: "root",
	}
	find(s, tree)
	dfs(tree)

	return nil
}
