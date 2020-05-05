package main

func romanToInt(s string) int {
	arrInt := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arrString := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	index := 0
	num := 0
	for index < 13 {
		//runtime.Breakpoint()
		for {
			if len(arrString[index]) == 1 {
				if len(s) < 1 {
					break
				}
				if s[0:1] == arrString[index] {
					num += arrInt[index]
					if len(s) == 1 {
						s = ""
					} else {
						s = s[1:]
					}
				} else {
					break
				}
			} else {
				if len(s) < 2 {
					break
				}
				if s[0:2] == arrString[index] {
					num += arrInt[index]
					if len(s) == 2 {
						s = ""
					} else {
						s = s[2:]
					}
				} else {
					break
				}
			}
		}
		index++
		if len(s) == 0 {
			break
		}
	}
	return num
}
