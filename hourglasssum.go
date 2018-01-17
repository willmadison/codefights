package codefights

import (
	"fmt"
	"math"
	"os"
)

const hourglassHeight = 3
const hourglassWidth = 3

func hourglassSum(a [][]int) int {
	maxSum := math.MinInt32
	var row int

	for row <= len(a)-hourglassHeight {
		var col int

		for col <= len(a)-hourglassWidth {
			sum := calculateHourglassSumAt(row, col, a)

			fmt.Fprintf(os.Stdout, "hourglassSumAt(%v, %v) = %v (currentMax = %v)\n", row, col, sum, maxSum)

			if sum > maxSum {
				maxSum = sum
			}

			col++
		}

		row++
	}

	return maxSum
}

func calculateHourglassSumAt(row, col int, a [][]int) int {
	return a[row][col] + a[row][col+1] + a[row][col+2] + a[row+1][col+1] + a[row+2][col] + a[row+2][col+1] + a[row+2][col+2]
}
