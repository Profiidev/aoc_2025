package internal

import (
	"log"
	"strconv"
	"strings"
)

type operator int

const (
	addition operator = iota
	multiplication
)

type problem struct {
	numbers  []int
	operator operator
}

func Day06(input string) {
	problems := []problem{}

	for line := range strings.SplitSeq(input, "\n") {
		idx := 0
		for column := range strings.SplitSeq(line, " ") {
			if column == "" {
				continue
			} else if column == "+" {
				problems[idx].operator = addition
				idx++
				continue
			} else if column == "*" {
				problems[idx].operator = multiplication
				idx++
				continue
			}

			num, err := strconv.Atoi(column)
			if err != nil {
				log.Panicf("Error: %s", err)
			}

			if idx >= len(problems) {
				problems = append(problems, problem{
					numbers:  []int{},
					operator: addition,
				})
			}

			problems[idx].numbers = append(problems[idx].numbers, num)
			idx++
		}
	}

	result := 0

	for _, p := range problems {
		subResult := 0
		switch p.operator {
		case addition:
			for _, n := range p.numbers {
				subResult += n
			}
		case multiplication:
			subResult = 1
			for _, n := range p.numbers {
				subResult *= n
			}
		}
		result += subResult
	}

	log.Printf("Result: %d", result)

	chars := [][]rune{}

	for line := range strings.SplitSeq(input, "\n") {
		row := []rune{}
		for char := range strings.SplitSeq(line, "") {
			row = append(row, []rune(char)[0])
		}

		chars = append(chars, row)
	}

	problem_idx := 0
	problems = []problem{}

	for i := 0; i < len(chars[0]); i++ {
		num := ""
		for j := 0; j < len(chars)-1; j++ {
			char := chars[j][i]
			if char == ' ' {
				continue
			}
			num += string(char)
		}

		if num == "" {
			problem_idx++
			continue
		}

		n, err := strconv.Atoi(num)
		if err != nil {
			log.Panicf("Error: %s", err)
		}

		if problem_idx >= len(problems) {
			problems = append(problems, problem{
				numbers:  []int{},
				operator: addition,
			})
		}

		problems[problem_idx].numbers = append(problems[problem_idx].numbers, n)
	}

	problem_idx = 0

	for i := 0; i < len(chars[0]); i++ {
		char := chars[len(chars)-1][i]
		if char == ' ' {
			continue
		}

		switch char {
		case '+':
			problems[problem_idx].operator = addition
			problem_idx++
		case '*':
			problems[problem_idx].operator = multiplication
			problem_idx++
		}
	}

	result = 0

	for _, p := range problems {
		subResult := 0
		switch p.operator {
		case addition:
			for _, n := range p.numbers {
				subResult += n
			}
		case multiplication:
			subResult = 1
			for _, n := range p.numbers {
				subResult *= n
			}
		}
		result += subResult
	}

	log.Printf("Result 2: %d", result)
}
