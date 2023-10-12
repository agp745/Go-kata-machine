package twocrystalballs

import "math"

func two_crystal_balls(breaks []bool) int {
	jump := math.Sqrt(float64(len(breaks)))
	var i int

	for i = 0; i < len(breaks); i += int(jump) {
		if breaks[i] {
			break
		}
	}

	i -= int(jump)
	for j := 0; j < int(jump) && i < len(breaks); j++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}
