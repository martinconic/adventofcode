package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	list := []string{}
	oxygen := [][]string{}
	co2 := [][]string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		list = strings.Split(s, "")

		oxygen = append(oxygen, list)
		co2 = append(co2, list)
	}

	for i := 0; i < len(oxygen[0]); i++ {
		var c0, c1 int
		var s0, s1 [][]string

		if len(oxygen) == 1 {
			break
		}

		for _, rep := range oxygen {
			r, _ := strconv.Atoi(rep[i])
			if r == 1 {
				c1++
				s1 = append(s1, rep)
			} else {
				c0++
				s0 = append(s0, rep)
			}
		}

		if c1 >= c0 {
			oxygen = s1
		} else {
			oxygen = s0
		}
	}

	for i := 0; i < len(co2[0]); i++ {
		var c0, c1 int
		var s0, s1 [][]string

		if len(co2) == 1 {
			break
		}

		for _, rep := range co2 {
			r, _ := strconv.Atoi(rep[i])
			if r == 1 {
				c1++
				s1 = append(s1, rep)
			} else {
				c0++
				s0 = append(s0, rep)
			}
		}

		if c1 >= c0 {
			co2 = s0
		} else {
			co2 = s1
		}

	}

	value1, _ := strconv.ParseInt(strings.Join(oxygen[0], ""), 2, 64)
	value2, _ := strconv.ParseInt(strings.Join(co2[0], ""), 2, 64)

	log.Println(value1, value2, value1*value2)
}
