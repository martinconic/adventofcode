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

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	text := scanner.Text()

	l := len(text)
	duplicate := 1
	for i := 0; i < l-14; i++ {
		s := text[i : i+14]
		for _, c := range s {
			count := strings.Count(s, string(c))
			if count > 1 {
				duplicate = count
			}
		}

		if duplicate == 1 {
			log.Println(i + 14)
			break
		}

		duplicate = 1

	}
	//	log.Println(text)

}
