package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pipe struct {
	in  string
	out string
}

var lp map[string]bool
var area [][]string

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}
	lp = map[string]bool{}
	//	pipes := map[string]pipe{}

	defer file.Close()
	area = [][]string{}
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

	//hack
	x := start[0]
	y := start[1] + 1
	count := 1
	pp := pipe{in: "N", out: "E"}
	lp[strconv.Itoa(start[0])+","+strconv.Itoa(start[1])] = true

	for area[x][y] != "S" {
		lp[strconv.Itoa(x)+","+strconv.Itoa(y)] = true
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
	count := 0

	for i := 0; i < len(area); i++ {
		for j := 0; j < len(area[0]); j++ {
			if _, ok := lp[strconv.Itoa(i)+","+strconv.Itoa(j)]; !ok {
				area[i][j] = "."
			}
		}
	}

	//	for _, v := range area {
	//		fmt.Println(v)
	//	}

	for i := 1; i < len(area); i++ {
		for j := 1; j < len(area[0]); j++ {
			if _, ok := lp[strconv.Itoa(i)+","+strconv.Itoa(j)]; !ok {
				c := 0
				stack := []string{}
				for k := 0; k < i; k++ {
					if area[k][j] == "-" || area[i][k] == "S" {
						c++
					}
					if area[k][j] == "F" || area[k][j] == "7" {
						stack = append(stack, area[k][j])
					}

					if area[k][j] == "J" && len(stack) > 0 && stack[len(stack)-1] == "F" {
						c++
						stack = stack[:len(stack)-1]
					}
					if area[k][j] == "L" && len(stack) > 0 && stack[len(stack)-1] == "7" {
						c++
						stack = stack[:len(stack)-1]
					}

				}
				stack = []string{}
				if c%2 == 0 {
					continue
				}
				c = 0

				for k := i + 1; k < len(area); k++ {
					if area[k][j] == "-" || area[i][k] == "S" {
						c++
					}
					if area[k][j] == "F" || area[k][j] == "7" {
						stack = append(stack, area[k][j])
					}

					if area[k][j] == "J" && len(stack) > 0 && stack[len(stack)-1] == "F" {
						c++
						stack = stack[:len(stack)-1]
					}
					if area[k][j] == "L" && len(stack) > 0 && stack[len(stack)-1] == "7" {
						c++
						stack = stack[:len(stack)-1]
					}

				}
				stack = []string{}
				if c%2 == 0 {
					continue
				}
				c = 0

				for k := 0; k < j; k++ {
					if area[i][k] == "|" || area[i][k] == "S" {
						c++
					}
					if area[i][k] == "F" || area[i][k] == "L" {
						stack = append(stack, area[i][k])
					}

					if area[i][k] == "J" && len(stack) > 0 && stack[len(stack)-1] == "F" {
						c++
						stack = stack[:len(stack)-1]
					}
					if area[i][k] == "7" && len(stack) > 0 && stack[len(stack)-1] == "L" {
						c++
						stack = stack[:len(stack)-1]
					}

				}
				stack = []string{}
				if c%2 == 0 {
					continue
				}
				c = 0

				for k := j + 1; k < len(area[0]); k++ {
					if area[i][k] == "|" || area[i][k] == "S" {
						c++
					}

					if area[i][k] == "F" || area[i][k] == "L" {
						stack = append(stack, area[i][k])
					}

					if area[i][k] == "J" && len(stack) > 0 && stack[len(stack)-1] == "F" {
						c++
						stack = stack[:len(stack)-1]
					}
					if area[i][k] == "7" && len(stack) > 0 && stack[len(stack)-1] == "L" {
						c++
						stack = stack[:len(stack)-1]
					}

				}
				if c%2 == 0 {
					continue
				}
				//	fmt.Println(i, j, area[i][j])
				count++

			}
		}
	}
	fmt.Println("Tiles: ", count)
}

func main() {
	p1()
	p2()
}
