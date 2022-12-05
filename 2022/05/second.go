package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//	stacks := [][]string{{"Z", "N"}, {"M", "C", "D"}, {"P"}}

	stacks := [][]string{{"R", "G", "J", "B", "T", "V", "Z"}, {"J", "R", "V", "L"}, {"S", "Q", "F"},
		{"Z", "H", "N", "L", "F", "V", "Q", "G"}, {"R", "Q", "T", "J", "C", "S", "M", "W"},
		{"S", "W", "T", "C", "H", "F"}, {"D", "Z", "C", "V", "F", "N", "J"},
		{"L", "G", "Z", "D", "W", "R", "F", "Q"}, {"J", "B", "W", "V", "P"}}

	for scanner.Scan() {
		text := strings.Fields(scanner.Text())

		move, _ := strconv.Atoi(text[1])
		f, _ := strconv.Atoi(text[3])
		t, _ := strconv.Atoi(text[5])

		from := stacks[f-1]
		to := stacks[t-1]
		last := len(from)
		to = append(to, from[last-move:]...)
		from = from[:last-move]

		stacks[f-1] = from
		stacks[t-1] = to

	}
	message := ""
	for i := 0; i < len(stacks); i++ {
		message += stacks[i][len(stacks[i])-1]
	}

	log.Println(message)

}
