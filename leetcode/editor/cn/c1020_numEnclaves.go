package main

func checkBoarder(x int, y int, lengthX int, lengthY int) bool {
	if x < 0 || x > lengthX-1 ||
		y < 0 || y > lengthY-1 {
		return false
	}
	return true
}

func findClosed(arrA [][]int, pCheckA *[][]int, postI int, postJ int, goOutNum int, pValue *[][]int) int {
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, 1, 0, -1}

	for k := 0; k < 4; k++ {
		x := dirX[k] + postI
		y := dirY[k] + postJ
		if checkBoarder(x, y, len(arrA), len(arrA[0])) {
			if (*pCheckA)[x][y] == 0 {
				(*pCheckA)[x][y] = 1
				if arrA[x][y] != 0 {
					(*pValue)[x][y] = 1
					goOutNum += 1
					goOutNum = findClosed(arrA, pCheckA, x, y, goOutNum, pValue)
				}
			}
		}
	}

	return goOutNum
}

//5.3 17:00
//ending 5.3 23:30 -- 扣去吃饭1:30分，耗时 5小时
func numEnclaves(A [][]int) int {
	//(i, j)
	if len(A) == 1 {
		return 0
	}
	checkA := make([][]int, len(A))
	valueA := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		checkA[i] = make([]int, len(A[0]))
		valueA[i] = make([]int, len(A[0]))
	}
	goOutNum := 0
	totalNum := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				totalNum += 1
			}
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			// 第一行和最后一行都是边界
			if i != 0 && i != len(A)-1 {
				// 不是第一行最后一行，跳过中间值
				if j != 0 && j != len(A[0])-1 {
					continue
				}
			}
			if checkA[i][j] == 0 {
				checkA[i][j] = 1
				if A[i][j] == 1 {

					goOutNum += 1
					valueA[i][j] = 1
					// four direction
					// Up
					// TODO  find next ,直到找不到1为止
					goOutNum = findClosed(A, &checkA, i, j, goOutNum, &valueA)

				}
			}
		}
	}
	return totalNum - goOutNum
}
