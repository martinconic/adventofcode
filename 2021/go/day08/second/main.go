package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

var count int

func main() {
	mo := [][]string{}
	mi := [][]string{}
	vmap := make(map[string]int)

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

	for _, o := range mo {
		for i, value := range o {
			o[i] = SortString(value)
		}
	}

	sevenkey := ""
	fourkey := ""
	for index, m := range mi {
		for i, value := range m {
			m[i] = SortString(value)
			length := len(value)
			if length == 2 {
				vmap[m[i]] = 1

			} else if length == 4 {
				vmap[m[i]] = 4
				fourkey = m[i]

			} else if length == 3 {
				vmap[m[i]] = 7
				sevenkey = m[i]

			} else if length == 7 {
				vmap[m[i]] = 8
			} else {
				vmap[m[i]] = -1
			}
		}

		for key, value := range vmap {
			if value == -1 {
				rem := key
				for _, letter := range sevenkey {
					rem = strings.ReplaceAll(rem, string(letter), "")
				}
				if len(rem) == 2 {
					vmap[key] = 3
				} else if len(rem) == 4 {
					vmap[key] = 6
				} else {
					isnine := true
					for _, letter := range fourkey {
						if !strings.Contains(key, string(letter)) {
							isnine = false
						}
					}
					if isnine {
						vmap[key] = 9
					}
				}
			}
		}

		for key, value := range vmap {
			if value == -1 {
				if len(key) == 6 {
					vmap[key] = 0
				} else {
					digit := key
					for _, letter := range fourkey {
						if !strings.Contains(digit, string(letter)) {
							digit += string(letter)
						}
					}

					if len(digit) == 7 {
						vmap[key] = 2
					} else {
						vmap[key] = 5
					}
				}
			}
		}

		// log.Println(m)
		// log.Println(vmap)

		count += 1000*vmap[mo[index][0]] + 100*vmap[mo[index][1]] + 10*vmap[mo[index][2]] + vmap[mo[index][3]]
	}

	log.Println(count)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
