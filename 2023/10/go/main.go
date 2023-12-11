package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type pipe struct {
	in  string
	out string
}

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	//	pipes := map[string]pipe{}

	defer file.Close()
	area := [][]string{}
	start := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if indx := strings.Index(line, "S"); indx >= 0 {
			start = append(start, len(area), indx)
		}

		values := strings.Split(line, "")
		area = append(area, values)
	}
	fmt.Println(start)
	fmt.Println(area[start[0]][start[1]])

	//hack
	x := start[0]
	y := start[1] + 1
	count := 1
	pp := pipe{in: "N", out: "E"}

	for area[x][y] != "S" {
		count++
		//	fmt.Println(x, y, area[x][y])
		if area[x][y] == "7" {
			if pp.out == "E" {
				pp = pipe{in: "W", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "W"}
				y--
			}
		} else if area[x][y] == "J" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "W"}
				y--
			} else {
				pp = pipe{in: "W", out: "N"}
				x--
			}

		} else if area[x][y] == "F" {
			if pp.out == "W" {
				pp = pipe{in: "E", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "E"}
				y++
			}

		} else if area[x][y] == "L" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "E"}
				y++
			} else {
				pp = pipe{in: "E", out: "N"}
				x--
			}
		} else if area[x][y] == "-" {
			if pp.out == "E" {
				pp = pipe{in: "W", out: "E"}
				y++
			} else {
				pp = pipe{in: "E", out: "W"}
				y--
			}
		} else if area[x][y] == "|" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "N"}
				x--
			}
		}

	}
	fmt.Println(count / 2)
}

func p2() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatalf("can not open file")
	}

	//	pipes := map[string]pipe{}

	defer file.Close()
	area := [][]string{}
	start := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if indx := strings.Index(line, "S"); indx >= 0 {
			start = append(start, len(area), indx)
		}

		values := strings.Split(line, "")
		area = append(area, values)
	}
	fmt.Println(start)
	fmt.Println(area[start[0]][start[1]])

	//hack
	x := start[0]
	y := start[1] + 1
	count := 1
	pp := pipe{in: "N", out: "E"}

	for area[x][y] != "S" {
		count++
		//	fmt.Println(x, y, area[x][y])
		if area[x][y] == "7" {
			if pp.out == "E" {
				pp = pipe{in: "W", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "W"}
				y--
			}
		} else if area[x][y] == "J" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "W"}
				y--
			} else {
				pp = pipe{in: "W", out: "N"}
				x--
			}

		} else if area[x][y] == "F" {
			if pp.out == "W" {
				pp = pipe{in: "E", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "E"}
				y++
			}

		} else if area[x][y] == "L" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "E"}
				y++
			} else {
				pp = pipe{in: "E", out: "N"}
				x--
			}
		} else if area[x][y] == "-" {
			if pp.out == "E" {
				pp = pipe{in: "W", out: "E"}
				y++
			} else {
				pp = pipe{in: "E", out: "W"}
				y--
			}
		} else if area[x][y] == "|" {
			if pp.out == "S" {
				pp = pipe{in: "N", out: "S"}
				x++
			} else {
				pp = pipe{in: "S", out: "N"}
				x--
			}
		}

	}
	fmt.Println(count / 2)
}

func main() {
	p2()
}
