package internal

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type direction int

const (
	Left direction = iota
	Right
)

type instruction struct {
	direction direction
	amount    int
}

func Day01(input string) {
	lines := strings.Split(input, "\n")
	instructions := []instruction{}
	for _, line := range lines {
		direction := Left
		if line[0] == 'R' {
			direction = Right
		}

		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Failed to parse line: %s", line)
		}

		inst := instruction{
			direction: direction,
			amount:    i,
		}

		instructions = append(instructions, inst)
	}

	pointer := 50
	count := 0
	count_2 := 0

	for _, inst := range instructions {
		rot := inst.amount
		if inst.direction == Left {
			rot *= -1
		}

		count_2 += int(math.Abs(math.Floor(float64(pointer+rot) / 100.0)))
		if pointer == 0 && rot < 0 {
			count_2 -= 1
		}

		pointer = (pointer + rot) % 100
		for pointer < 0 {
			pointer += 100
		}

		if pointer == 0 {
			count += 1
		}

		if pointer == 0 && rot < 0 {
			count_2 += 1
		}
	}

	log.Printf("Pointer was %d times at 0", count)
	log.Printf("Pointer crossed / was %d times at 0", count_2)
}
