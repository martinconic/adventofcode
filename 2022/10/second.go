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

	//	cycles := map[int]int{20: 20, 60: 60, 100: 100, 140: 140, 180: 180, 220: 220}
	position := 0
	register := 1

	sprite := map[int]int{0: 0, 1: 1, 2: 2}

	text := ""

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 1 {
			if sprite[0] == position || sprite[1] == position || sprite[2] == position {
				text += "#"
			} else {
				text += "."
			}
			position++

		} else if len(line) == 2 {
			for i := 0; i < 2; i++ {
				if sprite[0] == position || sprite[1] == position || sprite[2] == position {
					text += "#"
				} else {
					text += "."
				}
				if position == 39 {
					position = 0
					log.Println(text)
					text = ""
				} else {
					position++
				}
			}

			r, _ := strconv.Atoi(line[1])
			register += r
			sprite[0] = register - 1
			sprite[1] = register
			sprite[2] = register + 1

		}

		if position == 40 {
			position = 0
			log.Println(text)
			text = ""
		}
	}
}
