package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var count int

func main() {
	openings := "([{<"
	// closings := ")]}>"

	chunksmap := map[string]string{"{": "}", "(": ")", "[": "]", "<": ">"}

	openchunk := []string{}
	closechunk := []string{}

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
					calculate(value)
					break
				}
			}

			// log.Println(openchunk)
			// log.Println(closechunk)

		}

	}

	log.Println(count)

}

func calculate(value string) {
	if value == ")" {
		count += 3
	} else if value == "]" {
		count += 57
	} else if value == "}" {
		count += 1197
	} else if value == ">" {
		count += 25137
	}
}
