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

	sum := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")

		p10, _ := strconv.Atoi(p1[0])
		p11, _ := strconv.Atoi(p1[1])
		p20, _ := strconv.Atoi(p2[0])
		p21, _ := strconv.Atoi(p2[1])

		if (p10 >= p20 && p10 <= p21) || (p11 <= p21 && p11 >= p20) ||
			(p20 >= p10 && p20 <= p11) || (p21 <= p11 && p21 >= p10) ||
			(p10 <= p20 && p11 >= p21) || (p20 <= p10 && p21 >= p11) {

			sum++
		} else {

			log.Println(p2[1], "<=", p1[1], p2[1], ">=", p10, p21 <= p11, p21 >= p10)
		}
	}

	log.Println(sum)

}
