package main

func checkBoarder(x int, y int, lengthX int, lengthY int) bool {
	if x < 0 || x > lengthX-1 ||
		y < 0 || y > lengthY-1 {
		return false
	}
	return true
}

func findClosed(arrA *[][]int, postI int, postJ int, oldColor int, newColor int) {
	dirX := []int{-1, 0, 1, 0}
	dirY := []int{0, 1, 0, -1}

	for k := 0; k < 4; k++ {
		x := dirX[k] + postI
		y := dirY[k] + postJ
		if checkBoarder(x, y, len(*arrA), len((*arrA)[0])) {
			if (*arrA)[x][y] == oldColor {
				(*arrA)[x][y] = newColor
				findClosed(arrA, x, y, oldColor, newColor)
			}
		}
	}
}

// 23:40  ending 24:00
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	image[sr][sc] = newColor
	if oldColor == newColor {
		return image
	}
	findClosed(&image, sr, sc, oldColor, newColor)
	return image
}
