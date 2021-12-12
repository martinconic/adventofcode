package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var count int
var matrix [][]int
var cmap map[string][]string

func main() {

	matrix = [][]int{}
	cmap = make(map[string][]string)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "-")
		if _, ok := cmap[data[0]]; ok {
			cmap[data[0]] = append(cmap[data[0]], data[1])
		} else {
			cmap[data[0]] = []string{data[1]}
		}

		if _, ok := cmap[data[1]]; ok {
			cmap[data[1]] = append(cmap[data[1]], data[0])
		} else {
			cmap[data[1]] = []string{data[0]}
		}
	}
	log.Println(cmap)

	for _, val := range cmap["start"] {
		path := "start" + "," + val
		findpath(val, path)
	}

	log.Println(count)

}

func findpath(value, path string) {
	for _, val := range cmap[value] {
		if !strings.Contains(path, strings.ToLower(val)) {
			pth := path + "," + val
			if val == "end" {
				log.Println(pth)
				count++
			} else {
				findpath(val, pth)
			}
		}
	}
}
