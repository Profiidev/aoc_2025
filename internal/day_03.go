package internal

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Day03(input string) {
	banks := [][]int{}

	for line := range strings.SplitSeq(input, "\n") {
		bank := []int{}
		for digit := range strings.SplitSeq(line, "") {
			num, err := strconv.Atoi(digit)
			if err != nil {
				log.Panicf("error: %s", err)
			}
			bank = append(bank, num)
		}
		banks = append(banks, bank)
	}

	result := 0
	result_2 := 0

	for _, bank := range banks {
		highest := -1

		for i := 0; i < len(bank)-1; i++ {
			for j := i + 1; j < len(bank); j++ {
				num := bank[i]*10 + bank[j]
				if num > highest {
					highest = num
				}
			}
		}

		result += highest
		result_2 += findHighestInPart(bank, 12)
	}

	log.Printf("Result: %d", result)
	log.Printf("Result 2: %d", result_2)
}

func findHighestInPart(part []int, length int) int {
	if length == 0 {
		return 0
	}

	highest := -1
	for i := 9; i > 0; i-- {
		for j, digit := range part {
			if digit == i {
				sub_part := findHighestInPart(part[j+1:], length-1)
				if sub_part == -1 {
					continue
				}
				num := digit*int(math.Pow(10, float64(length)-1)) + sub_part
				if num > highest {
					highest = num
				}
			}
		}

		if highest != -1 {
			return highest
		}
	}
	return -1
}
