package codefights

import "sort"

func avoidObstacles(obstacles []int) int {
	sort.Ints(obstacles)

	jumpSize := 2

	collides := true

	for collides {
		for _, obstacle := range obstacles {
			collides = obstacle%jumpSize == 0
			if collides {
				jumpSize++
				break
			}
		}
	}

	return jumpSize
}
