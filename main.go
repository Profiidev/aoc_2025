package main

import (
	"aoc_2025/internal"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Please provide a day")
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal("Invalid int provided for day")
	}

	day_file := fmt.Sprintf("day_%02d.txt", i)
	input_path := filepath.Join("input", day_file)
	data, err := os.ReadFile(input_path)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	input := string(data)

	switch i {
	case 1:
		internal.Day01(input)
	case 2:
		internal.Day02(input)
	case 3:
		internal.Day03(input)
	case 4:
		internal.Day04(input)
	case 5:
		internal.Day05(input)
	default:
		log.Fatalf("Day %d not found", i)
	}
}
