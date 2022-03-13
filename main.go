package main

import (
	"fmt"
)

func main() {
	// Test input values
	tVals := []int{-1, 0, 01, 100, 123, 121, 12344321, 4444, 101}
	for _, v := range tVals {
		fmt.Printf("\nInput: %d\n", v)
		fmt.Printf("Solution: %t\n", isPalindrome(v))
	}
}

func isPalindrome(x int) bool {
	x, y, chk := x, 0, false
	tmpSx, tmpSy := []int{}, []int{}

	// return false, if negative integer
	if x < 0 {
		return chk
	}

	// Loop over the int to build slices
	for i := 1; i <= x; i++ {
		y, x = x%10, x/10
		tmpSx = append(tmpSx, y)
	}

	if x != 0 {
		tmpSx = append(tmpSx, x)
	}
	tmpSy = append(tmpSy, tmpSx...)

	// Reverse one of our slices (tmpSy)
	for i, j := 0, len(tmpSy)-1; i < j; i, j = i+1, j-1 {
		tmpSy[i], tmpSy[j] = tmpSy[j], tmpSy[i]
	}

	// Ensure slices lengths match
	if len(tmpSx) == len(tmpSy) {
		for i, v := range tmpSx {
			if v != tmpSy[i] {
				return chk
			}
		}
		chk = true
	}
	return chk
}
