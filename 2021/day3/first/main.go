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
	list := make([]int, 12)
	var gl, el string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytesa := strings.Split(scanner.Text(), "")
		// log.Println(bytesa)
		for i, b := range bytesa {
			numb, _ := strconv.Atoi(b)
			if numb == 1 {
				list[i] += 1
			} else {
				list[i] -= 1
			}
		}
	}

	for _, g := range list {
		if g > 0 {
			gl += "1"
			el += "0"
		} else {
			gl += "0"
			el += "1"

		}
	}

	gamma, _ := strconv.ParseInt(gl, 2, 64)
	epsilon, _ := strconv.ParseInt(el, 2, 64)

	log.Println(gamma, epsilon, gamma*epsilon)

}
