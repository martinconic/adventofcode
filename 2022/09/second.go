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

	tail := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	n := len(tail)
	xpt := 0
	ypt := 0

	m := map[string]int{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		dir := line[0]
		steps, _ := strconv.Atoi(line[1])

		//log.Println("---", dir, steps, "---")

		s := "0,0"
		m[s] += 1

		for i := 0; i < steps; i++ {
			if dir == "R" {
				xh++
				if xh > xt && xh-xt == 2 {
					xt++

					if yh != yt {
						yt = yh
					}

				}
			}

			if dir == "L" {
				xh--
				if xh < xt && xt-xh == 2 {
					xt--

					if yh != yt {
						yt = yh
					}

				}
			}

			if dir == "U" {
				yh++
				if yh > yt && yh-yt == 2 {
					yt++

					if xh != xt {
						xt = xh
					}

				}
			}

			if dir == "D" {
				yh--
				if yh < yt && yt-yh == 2 {
					yt--
					if xh != xt {
						xt = xh
					}

				}
			}

			//log.Println(xh, yh, xt, yt)

			if xt != tail[0][0] || yt != tail[0][1] {
				x := xt
				y := yt
				for j := 0; j < n; j++ {
					if x > tail[j][0] && x-tail[j][0] == 2 {
						tail[j][0] += 1
						if y > tail[j][1] {
							tail[j][1] += 1
						}
						if y < tail[j][1] {
							tail[j][1] -= 1
						}

					} else if x < tail[j][0] && tail[j][0]-x == 2 {
						tail[j][0] -= 1
						if y > tail[j][1] {
							tail[j][1] += 1
						} else if y < tail[j][1] {
							tail[j][1] -= 1
						}

					} else if y > tail[j][1] && y-tail[j][1] == 2 {
						tail[j][1] += 1
						if x > tail[j][0] {
							tail[j][0] += 1
						} else if x < tail[j][0] {
							tail[j][0] -= 1
						}

					} else if y < tail[j][1] && tail[j][1]-y == 2 {
						tail[j][1] -= 1

						if x > tail[j][0] {
							tail[j][0] += 1
						} else if x < tail[j][0] {
							tail[j][0] -= 1
						}

					}

					x = tail[j][0]
					y = tail[j][1]
				}

			}

			if tail[n-1][0] != xpt || tail[n-1][1] != ypt {
				xpt = tail[n-1][0]
				ypt = tail[n-1][1]
				s = strconv.Itoa(xpt) + "," + strconv.Itoa(ypt)
				m[s] += 1
			}

		}

	}

	log.Println(len(m))
}
