package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	elem := make(map[string]bool)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	template := scanner.Text()

	rules := make(map[string]string)
	scanner.Scan()
	for scanner.Scan() {
		str := scanner.Text()
		data := strings.Split(str, " -> ")
		rules[data[0]] = data[1]
		if _, ok := elem[data[1]]; !ok {
			elem[data[1]] = true
		}
	}

	length := len(template)
	for k := 0; k < 40; k++ {
		for i := 0; i < length-1; i++ {
			if value, ok := rules[template[i:i+2]]; ok {
				template = template[:i+1] + value + template[i+1:]
				length++
				i++
			}
		}
		// log.Println(template)

	}

	max := 0
	min := 10000

	for key := range elem {
		count := strings.Count(template, key)

		if count > max {
			max = count
		}

		if count < min {
			min = count
		}
	}
	log.Println(max)
	log.Println(min)

	log.Println(max - min)
}
