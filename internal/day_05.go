package internal

import (
	"log"
	"strconv"
	"strings"
)

type freshRange struct {
	start int
	end   int
}

func (rang *freshRange) contains(pos int) bool {
	return pos >= rang.start && pos <= rang.end
}

func Day05(input string) {
	ranges := []freshRange{}
	ids := []int{}
	in_ranges := true

	for line := range strings.SplitSeq(input, "\n") {
		if line == "" {
			in_ranges = false
			continue
		}

		if in_ranges {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Panicf("Error: %s", err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Panicf("Error: %s", err)
			}

			ranges = append(ranges, freshRange{
				start: start,
				end:   end,
			})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				log.Panicf("Error: %s", err)
			}
			ids = append(ids, id)
		}
	}

	result := 0

	for _, id := range ids {
		for _, fresh_range := range ranges {
			if fresh_range.contains(id) {
				result++
				break
			}
		}
	}

	result_2 := 0
	completed := []freshRange{}
	for _, fresh_range := range ranges {
		todo_ranges := []freshRange{fresh_range}
		done_ranges := []freshRange{}
	top:
		for {
			if len(todo_ranges) == 0 {
				break
			}

			todo := todo_ranges[0]
			todo_ranges = todo_ranges[1:]

			for _, complete := range completed {
				if complete.end < todo.start || complete.start > todo.end {
					continue
				}

				intersected := false
				if todo.contains(complete.start) {
					if complete.start != todo.start {
						todo_ranges = append(todo_ranges, freshRange{
							start: todo.start,
							end:   complete.start - 1,
						})
					}
					intersected = true
				}

				if todo.contains(complete.end) {
					if complete.end != todo.end {
						todo_ranges = append(todo_ranges, freshRange{
							start: complete.end + 1,
							end:   todo.end,
						})
					}
					intersected = true
				}

				if complete.contains(todo.start) && complete.contains(todo.end) {
					intersected = true
				}

				if intersected {
					continue top
				}
			}

			done_ranges = append(done_ranges, todo)
		}

		for _, done := range done_ranges {
			completed = append(completed, done)
			result_2 += done.end - done.start + 1
		}
	}

	log.Printf("Result: %d", result)
	log.Printf("Result 2: %d", result_2)
}
