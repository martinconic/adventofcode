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

	xt := 0
	yt := 0

	xh := 0
	yh := 0

	m := map[string]int{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		dir := line[0]
		steps, _ := strconv.Atoi(line[1])

		//log.Println("---", dir, steps, "---")

		s := "0,0"
		m[s] += 1

		if dir == "R" {
			for i := 0; i < steps; i++ {
				xh++
				if xh > xt && xh-xt == 2 {
					xt++

					if yh != yt {
						yt = yh
					}

					s = strconv.Itoa(xt) + "," + strconv.Itoa(yt)
					m[s] += 1
				}
			}
		}

		if dir == "L" {
			for i := 0; i < steps; i++ {
				xh--
				if xh < xt && xt-xh == 2 {
					xt--

					if yh != yt {
						yt = yh
					}

					s = strconv.Itoa(xt) + "," + strconv.Itoa(yt)
					m[s] += 1

				}
			}
		}

		if dir == "U" {
			for i := 0; i < steps; i++ {
				yh++
				if yh > yt && yh-yt == 2 {
					yt++

					if xh != xt {
						xt = xh
					}

					s = strconv.Itoa(xt) + "," + strconv.Itoa(yt)
					m[s] += 1

				}
			}
		}

		if dir == "D" {
			for i := 0; i < steps; i++ {
				yh--
				if yh < yt && yt-yh == 2 {
					yt--

					if xh != xt {
						xt = xh
					}

					s = strconv.Itoa(xt) + "," + strconv.Itoa(yt)
					m[s] += 1

				}
			}
		}

	}

	log.Println(len(m))
}
