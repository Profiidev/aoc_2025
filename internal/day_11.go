package internal

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

type matrix = [][]bool

type data struct {
	matrix matrix
	nodes  []string
}

func (d *data) conn(node string) []string {
	i := slices.Index(d.nodes, node)
	if i == -1 {
		return []string{}
	}

	ret := []string{}
	for i, conn := range d.matrix[i] {
		if conn {
			ret = append(ret, d.nodes[i])
		}
	}
	return ret
}

func Day11(input string) {
	machines := make(map[string][]string)
	unique_map := make(map[string]bool)

	for line := range strings.SplitSeq(input, "\n") {
		parts := strings.Split(line, ": ")
		output := strings.Split(parts[1], " ")

		machines[parts[0]] = output

		unique_map[parts[0]] = true
		for _, out := range output {
			unique_map[out] = true
		}
	}

	unique := make([]string, len(unique_map))
	i := 0

	for k := range unique_map {
		unique[i] = k
		i++
	}
	slices.Sort(unique)

	connections := matrix{}
	for range len(unique) {
		sub := []bool{}
		for range len(unique) {
			sub = append(sub, false)
		}
		connections = append(connections, sub)
	}

	for i, a := range unique {
		outputs := machines[a]
		for _, out := range outputs {
			index := slices.Index(unique, out)
			if index != -1 {
				connections[i][index] = true
			}
		}
	}

	d := data{
		matrix: connections,
		nodes:  unique,
	}

	result := walk(d, "you", false)
	log.Printf("Result: %d", result)

	result_2 := walk(d, "svr", true)
	log.Printf("Result 2: %d", result_2)
}

func walk(d data, node string, check bool) int {
	return walk_inner(d, node, []string{node}, make(map[string]int), check)
}

func walk_inner(d data, node string, visited []string, found map[string]int, check bool) int {
	if node == "out" {
		valid := true
		for _, req := range []string{"dac", "fft"} {
			if !slices.Contains(visited, req) {
				valid = false
			}
		}

		if valid || !check {
			return 1
		} else {
			return 0
		}
	}

	key := fmt.Sprintf("%s%t%t", node, slices.Contains(visited, "fft"), slices.Contains(visited, "dac"))

	exists := found[key]
	if exists != 0 {
		if exists == -1 {
			return 0
		} else {
			return exists
		}
	}

	out := 0
	for _, next := range d.conn(node) {
		out += walk_inner(d, next, append(visited, next), found, check)
	}

	if out == 0 {
		found[key] = -1
	} else {
		found[key] = out
	}

	return out
}
