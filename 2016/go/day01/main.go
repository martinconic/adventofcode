package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func first(data []string) {
	up := true
	dx := true
	dist := []int{0, 0}
	for _, val := range data {
		v := strings.SplitN(val, "", 2)
		d, _ := strconv.Atoi(v[1])

		if dx {
			if v[0] == "R" {
				if up {
					dist[0] += d
					up = true
				} else {
					dist[0] -= d
					up = false
				}
			} else {
				if up {
					dist[0] -= d
					up = false
				} else {
					dist[0] += d
					up = true
				}
			}

			dx = false
		} else {
			if v[0] == "L" {
				if up {
					dist[1] += d
					up = true
				} else {
					dist[1] -= d
					up = false
				}
			} else {
				if up {
					dist[1] -= d
					up = false
				} else {
					dist[1] += d
					up = true
				}
			}
			dx = true
		}
		log.Println(dist)
	}

}

func second(data []string) {
	up := true
	dx := true
	dist := []int{0, 0}
	m := map[string]bool{}
	for _, val := range data {
		v := strings.SplitN(val, "", 2)
		d, _ := strconv.Atoi(v[1])

		if dx {
			if v[0] == "R" {
				if up {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0] + i), fmt.Sprint(dist[1])}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[0] += d
					up = true
				} else {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0] - i), fmt.Sprint(dist[1])}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[0] -= d
					up = false
				}
			} else {
				if up {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0] - i), fmt.Sprint(dist[1])}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[0] -= d
					up = false
				} else {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0] + i), fmt.Sprint(dist[1])}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[0] += d
					up = true
				}
			}

			dx = false
		} else {
			if v[0] == "L" {
				if up {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0]), fmt.Sprint(dist[1] + i)}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[1] += d
					up = true
				} else {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0]), fmt.Sprint(dist[1] - i)}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[1] -= d
					up = false
				}
			} else {
				if up {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0]), fmt.Sprint(dist[1] - i)}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[1] -= d
					up = false
				} else {
					for i := 0; i < d; i++ {
						s := []string{fmt.Sprint(dist[0]), fmt.Sprint(dist[1] + i)}
						ss := strings.Join(s, ",")
						if _, ok := m[ss]; ok {
							fmt.Println(ss)
							return
						}
						m[ss] = true
					}
					dist[1] += d
					up = true
				}
			}
			dx = true
		}
		log.Println(dist)
	}

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ", ")
	}

	log.Println(data)

	// first(data)
	second(data)
}
