package main

import "math"

func maxArea2(height []int) int {
	max := 0
	for x := 0; x < len(height)-1; x++ {
		for y := x + 1; y < len(height); y++ {
			area := 0
			if height[x] > height[y] {
				area = (y - x) * height[y]
			} else {
				area = (y - x) * height[x]
			}
			if area > max {
				max = area
			}
		}
	}
	return max
}

func maxArea(height []int) int {
	// 双指针。每次较短到一块板向另一个指针移动,此时指针距离 - 1
	//原因, 较长的板移动, 由于短板值更小,因此,移动后到结果一定是小于或等于短板的
	// 而移动短板,由于较短一方发生变化,因此短板可能变长,因此总面积就可能变大
	x := 0
	y := len(height) - 1
	max := 0
	for x != y {
		area := int(math.Min(float64(height[x]), float64(height[y]))) * (y - x)
		if area > max {
			max = area
		}
		if height[x] > height[y] {
			y = y - 1
		} else {
			x = x + 1
		}
	}
	return max
}
