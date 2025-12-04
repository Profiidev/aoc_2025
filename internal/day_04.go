package internal

import (
	"log"
	"strings"
)

func Day04(input string) {
	roles := [][]bool{}

	for line := range strings.SplitSeq(input, "\n") {
		out_line := make([]bool, len(line))
		for j, cell := range strings.Split(line, "") {
			out_line[j] = cell == "@"
		}
		roles = append(roles, out_line)
	}

	roles, result := RemoveRoles(roles)
	log.Printf("Result: %d", result)

	removed := 0
	for {
		roles, removed = RemoveRoles(roles)
		if removed == 0 {
			break
		}

		result += removed
	}

	log.Printf("Result 2: %d", result)
}

func RemoveRoles(roles [][]bool) ([][]bool, int) {
	result := 0
	roles_out := make([][]bool, len(roles))
	for i := range roles {
		roles_out[i] = make([]bool, len(roles[i]))
		copy(roles_out[i], roles[i])
	}

	for i := range roles {
		for j := 0; j < len(roles[0]); j++ {
			next := 0

			if !roles[i][j] {
				continue
			}

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}

					ix := i + x
					jy := j + y
					if ix < 0 || jy < 0 || ix >= len(roles) || jy >= len(roles[0]) {
						continue
					}

					if roles[ix][jy] {
						next++
					}
				}
			}

			if next < 4 {
				result++
				roles_out[i][j] = false
			}
		}
	}

	return roles_out, result
}
