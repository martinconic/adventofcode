package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

var count int
var prevcount int

func main() {
	mo := [][]string{}
	mi := [][]string{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	f := func(c rune) bool {
		return unicode.IsSpace(c) || c == '.'
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		data := strings.Split(scanner.Text(), "|")

		input := strings.FieldsFunc(data[0], f)
		output := strings.FieldsFunc(data[1], f)

		mi = append(mi, input)
		mo = append(mo, output)

	}

	count := 0
	for _, m := range mo {
		for _, value := range m {
			length := len(value)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}
	log.Println(count)
}
