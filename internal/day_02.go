package internal

import (
	"log"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

func Day02(input string) {
	ranges := []idRange{}

	for raw_range := range strings.SplitSeq(input, ",") {
		pair := strings.Split(raw_range, "-")
		start, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Panicf("Failed to extract number: %s", err)
		}
		end, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Panicf("Failed to extract number: %s", err)
		}

		ranges = append(ranges, idRange{
			start: start, end: end,
		})
	}

	result := 0

	for _, idRange := range ranges {
		for i := idRange.start; i <= idRange.end; i++ {
			i_str := strconv.Itoa(i)
			i_len := len(i_str)
			if i_len%2 == 0 && i_str[i_len/2:] == i_str[:i_len/2] {
				result += i
			}
		}
	}

	log.Printf("Result: %d", result)

	result = 0

	for _, idRange := range ranges {
		for i := idRange.start; i <= idRange.end; i++ {
			i_str := strconv.Itoa(i)
			if isRepeated(i_str) {
				result += i
			}
		}
	}

	log.Printf("Result: %d", result)
}

func isRepeated(s string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	for l := 1; l <= n/2; l++ {
		if n%l == 0 {
			substring := s[:l]
			isRep := true
			for i := l; i < n; i += l {
				if s[i:i+l] != substring {
					isRep = false
					break
				}
			}
			if isRep {
				return true
			}
		}
	}
	return false
}
