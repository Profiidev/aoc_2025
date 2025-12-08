package internal

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type vec3 struct {
	x int
	y int
	z int
}

type conn struct {
	a    vec3
	b    vec3
	dist float64
}

func (a *vec3) dist(b vec3) float64 {
	part_x := math.Pow(math.Abs(float64(a.x-b.x)), 2)
	part_y := math.Pow(math.Abs(float64(a.y-b.y)), 2)
	part_z := math.Pow(math.Abs(float64(a.z-b.z)), 2)
	return math.Sqrt(part_x + part_y + part_z)
}

func Day08(input string) {
	positions := []vec3{}

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
		z, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Panicf("Error: %s", err)
		}

		positions = append(positions, vec3{
			x: x,
			y: y,
			z: z,
		})
	}

	distances := []conn{}
	for i, a := range positions {
		for _, b := range positions[i+1:] {
			distances = append(distances, conn{
				a:    a,
				b:    b,
				dist: a.dist(b),
			})
		}
	}
	slices.SortFunc(distances, func(a conn, b conn) int {
		return int(a.dist - b.dist)
	})

	connections := make(map[vec3][]vec3)
	circuits := [][]vec3{}

	for i := 0; true; i++ {
		conn := distances[i]
		short_a := conn.a
		short_b := conn.b

		circuit_a := -1
		circuit_b := -1
		for i, circuit := range circuits {
			if slices.Contains(circuit, short_a) {
				circuit_a = i
			}
			if slices.Contains(circuit, short_b) {
				circuit_b = i
			}
		}

		connections[short_a] = append(connections[short_a], short_b)
		connections[short_b] = append(connections[short_b], short_a)

		if circuit_a != -1 && circuit_a == circuit_b {
			continue
		}

		if circuit_a == -1 && circuit_b == -1 {
			circuits = append(circuits, []vec3{short_a, short_b})
		} else if circuit_a == -1 {
			circuits[circuit_b] = append(circuits[circuit_b], short_a)
		} else if circuit_b == -1 {
			circuits[circuit_a] = append(circuits[circuit_a], short_b)
		} else {
			circuits[circuit_a] = append(circuits[circuit_a], circuits[circuit_b]...)
			circuits = append(circuits[:circuit_b], circuits[circuit_b+1:]...)
		}

		if i == 999 {
			longest := []int{}
			for _, circuit := range circuits {
				longest = append(longest, len(circuit))
			}
			slices.Sort(longest)

			result := 1
			for i := 1; i <= 3; i++ {
				result *= longest[len(longest)-i]
			}

			log.Printf("Result: %d", result)
		}

		if len(circuits) > 0 && len(circuits[0]) == len(positions) {
			log.Printf("Result 2: %d", short_a.x*short_b.x)
			break
		}
	}
}
