package internal

import (
	"log"
	"strconv"
	"strings"
)

type present struct {
	shape [][]bool
	area  int
	index int
}

type area struct {
	width    int
	height   int
	required map[int]int
}

func Day12(input string) {
	presents := []present{}
	areas := []area{}
	presentIdx := -1

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			continue
		}

		if line[0] >= '0' && line[0] <= '9' && line[1] == ':' {
			presentIdx++
			presents = append(presents, present{
				area:  0,
				shape: [][]bool{},
				index: presentIdx,
			})
		} else if line[0] == '#' || line[0] == '.' {
			row := []bool{}
			for _, char := range line {
				if char == '.' {
					row = append(row, false)
				} else {
					presents[presentIdx].area++
					row = append(row, true)
				}
			}
			presents[presentIdx].shape = append(presents[presentIdx].shape, row)
		} else {
			parts := strings.Split(line, ": ")
			dim := strings.Split(parts[0], "x")
			width, err := strconv.Atoi(dim[0])
			if err != nil {
				log.Panicf("Error: %s", err)
			}
			height, err := strconv.Atoi(dim[1])
			if err != nil {
				log.Panicf("Error: %s", err)
			}

			required := make(map[int]int)
			for i, num := range strings.Split(parts[1], " ") {
				num, err := strconv.Atoi(num)
				if err != nil {
					log.Panicf("Error: %s", err)
				}
				required[i] = num
			}

			areas = append(areas, area{
				width:    width,
				height:   height,
				required: required,
			})
		}
	}

	result := 0

	for _, area := range areas {
		total := 0
		for i, count := range area.required {
			total += count * presents[i].area
		}

		if total <= area.height*area.width {
			result++
		}
	}

	log.Printf("Result: %d", result)
}
