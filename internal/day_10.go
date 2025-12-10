package internal

import (
	"cmp"
	"container/heap"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

type machine struct {
	lamps_current  string
	lamps_required string
	buttons        [][]int
	joltage        []int
}

type state struct {
	current any
	toggled int
}

type queue []*state

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	return q[i].toggled < q[j].toggled
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *queue) Fix() {
	slices.SortFunc(*q, func(a, b *state) int {
		return cmp.Compare(b.toggled, a.toggled)
	})
}

func (q *queue) Push(x any) {
	*q = append(*q, x.(*state))
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}

func Day10(input string) {
	machines := []machine{}

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, " ")

		lamps_current := ""
		lamps_required := parts[0][1 : len(parts[0])-1]
		for i := 0; i < len(lamps_required); i++ {
			lamps_current += "."
		}

		buttons := [][]int{}
		for _, button_raw := range parts[1 : len(parts)-1] {
			button := []int{}
			for num_raw := range strings.SplitSeq(button_raw[1:len(button_raw)-1], ",") {
				num, err := strconv.Atoi(num_raw)
				if err != nil {
					log.Panicf("Error: %s", err)
				}
				button = append(button, num)
			}
			buttons = append(buttons, button)
		}

		jolatge := []int{}
		last := parts[len(parts)-1]
		for num_raw := range strings.SplitSeq(last[1:len(last)-1], ",") {
			num, err := strconv.Atoi(num_raw)
			if err != nil {
				log.Panicf("Error: %s", err)
			}
			jolatge = append(jolatge, num)
		}

		machines = append(machines, machine{
			lamps_current:  lamps_current,
			lamps_required: lamps_required,
			buttons:        buttons,
			joltage:        jolatge,
		})
	}

	result := 0

machine:
	for _, machine := range machines {
		q := &queue{
			&state{
				current: machine.lamps_current,
				toggled: 0,
			},
		}
		heap.Init(q)
		seen := make(map[string]bool)

		for q.Len() > 0 {
			current := heap.Pop(q).(*state)
			cur_state := current.current.(string)
			for _, btn := range machine.buttons {
				next := []rune(cur_state)
				for _, i := range btn {
					if next[i] == '.' {
						next[i] = '#'
					} else {
						next[i] = '.'
					}
				}
				if string(next) == machine.lamps_required {
					result += current.toggled + 1
					continue machine
				}

				if !seen[string(next)] {
					seen[cur_state] = true
					heap.Push(q, &state{
						current: string(next),
						toggled: current.toggled + 1,
					})
				}
			}
		}
	}

	log.Printf("Result: %d", result)

	result_2 := 0
	const maxPresses = 1000

	for _, machine := range machines {
		joltage := make([]int, len(machine.joltage))
		if slices.Equal(joltage, machine.joltage) {
			continue
		}

		numBtn := len(machine.buttons)
		numJol := len(machine.joltage)

		lp := golp.NewLP(0, numBtn)
		lp.SetVerboseLevel(golp.NEUTRAL)

		objCoeffs := make([]float64, numBtn)
		for i := range numBtn {
			objCoeffs[i] = 1.0
		}
		lp.SetObjFn(objCoeffs)

		for i := range numBtn {
			lp.SetInt(i, true)
			lp.SetBounds(i, 0.0, float64(maxPresses))
		}

		for i := 0; i < numJol; i++ {
			var entries []golp.Entry
			for j, btn := range machine.buttons {
				if slices.Contains(btn, i) {
					entries = append(entries, golp.Entry{Col: j, Val: 1.0})
				}
			}
			targetValue := float64(machine.joltage[i])
			if err := lp.AddConstraintSparse(entries, golp.EQ, targetValue); err != nil {
				log.Panicf("Error: %s", err)
			}
		}

		status := lp.Solve()

		if status != golp.OPTIMAL {
			continue
		}

		solution := lp.Variables()
		for _, val := range solution {
			result_2 += int(val + 0.5)
		}
	}

	log.Printf("Result 2: %d", result_2)
}
