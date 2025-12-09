package internal

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func Day09(input string) {
	points := []point{}

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ",")

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Panicf("Error: %s", err)
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Panicf("Error: %s", err)
		}

		points = append(points, point{
			x: x,
			y: y,
		})
	}

	result := 0

	for _, a := range points {
		for _, b := range points {
			area := int(math.Abs(float64(a.x-b.x)+1.0) * math.Abs(float64(a.y-b.y)+1.0))
			if area > result {
				result = area
			}
		}
	}

	greenPoints := []point{}
	for i := range len(points) - 1 {
		a := points[i]
		b := points[i+1]
		greenPoints = append(greenPoints, addPoints(a, b)...)
	}
	greenPoints = append(greenPoints, addPoints(points[len(points)-1], points[0])...)

	result_2 := 0
	for _, a := range points {
	inner:
		for _, b := range points {
			area := int(math.Abs(float64(a.x-b.x)+1.0) * math.Abs(float64(a.y-b.y)+1.0))
			if area < result_2 {
				continue
			}
			for _, green := range greenPoints {
				if green.x > min(a.x, b.x) && green.x < max(a.x, b.x) && green.y > min(a.y, b.y) && green.y < max(a.y, b.y) {
					continue inner
				}
			}
			result_2 = area
		}
	}

	log.Printf("Result: %d", result)
	log.Printf("Result 2: %d", result_2)
}

func addPoints(a point, b point) []point {
	greenPoints := []point{}

	if a.x == b.x {
		for y := min(a.y, b.y); y < max(a.y, b.y); y++ {
			greenPoints = append(greenPoints, point{
				x: a.x,
				y: y,
			})
		}
	} else if a.y == b.y {
		for x := min(a.x, b.x); x < max(a.x, b.x); x++ {
			greenPoints = append(greenPoints, point{
				x: x,
				y: a.y,
			})
		}
	} else {
		log.Panicf("Unaligned")
	}

	return greenPoints
}
