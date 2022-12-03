package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	items := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	scanner := bufio.NewScanner(file)
	sum := 0
	rs := []string{}
	index := 1

	for scanner.Scan() {
		rs = append(rs, scanner.Text())
		if index > 0 && index%3 == 0 {
			//log.Println(rs)
			for _, c := range items {
				if strings.Contains(rs[0], string(c)) && strings.Contains(rs[1], string(c)) &&
					strings.Contains(rs[2], string(c)) {
					sum += strings.Index(items, string(c)) + 1
				}
			}
			index = 0
			rs = []string{}
		}
		index++
	}

	log.Println(sum)

}
