package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	openings := "([{<"
	// closings := ")]}>"

	chunksmap := map[string]string{"{": "}", "(": ")", "[": "]", "<": ">"}

	openchunk := []string{}
	closechunk := []string{}
	scores := []int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")

		for _, value := range data {
			if strings.Contains(openings, value) {
				openchunk = append(openchunk, value)
				closechunk = append(closechunk, chunksmap[value])
			} else {
				if closechunk[len(closechunk)-1] == value {
					closechunk = closechunk[:len(closechunk)-1]
					openchunk = openchunk[:len(openchunk)-1]
				} else {
					openchunk = []string{}
					closechunk = []string{}
					break
				}
			}
		}

		if len(closechunk) > 0 {
			score := calculate(closechunk)
			scores = append(scores, score)
			log.Println(closechunk, score)
		}

		openchunk = []string{}
		closechunk = []string{}
	}

	sort.Ints(scores)

	index := (len(scores) - 1) / 2
	log.Println(scores[index])

}

func calculate(value []string) int {
	score := 0
	for i := len(value) - 1; i >= 0; i-- {
		score *= 5
		if value[i] == ")" {
			score += 1
		} else if value[i] == "]" {
			score += 2
		} else if value[i] == "}" {
			score += 3
		} else if value[i] == ">" {
			score += 4
		}
	}

	return score
}
