package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cards = map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

var cards2 = map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 10, "J": 0, "Q": 12, "K": 13, "A": 14}

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	hands := map[string]int{}

	fivek := []string{}
	fourk := []string{}
	fullh := []string{}
	threek := []string{}
	twop := []string{}
	onep := []string{}
	highc := []string{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		key := line[0]
		value, _ := strconv.Atoi(line[1])
		hands[key] = value

		hand := strings.Split(key, "")
		mhand := map[string]int{}

		for _, v := range hand {
			mhand[v] += 1
		}

		if len(mhand) == 1 {
			fivek = append(fivek, key)
			continue
		}
		if len(mhand) == 5 {
			highc = append(highc, key)
			continue
		}
		if len(mhand) == 4 {
			onep = append(onep, key)
			continue
		}
		if len(mhand) == 2 {
			if mhand[string(key[0])] == 1 || mhand[string(key[0])] == 4 {
				fourk = append(fourk, key)
			} else {
				fullh = append(fullh, key)
			}
			continue
		}
		if len(mhand) == 3 {
			if mhand[string(key[0])] == 2 || mhand[string(key[1])] == 2 {
				twop = append(twop, key)
			} else {
				threek = append(threek, key)
			}
			continue
		}

	}

	slices.SortFunc(fivek, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(fourk, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(fullh, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(threek, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(twop, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(onep, func(a, b string) int {
		return sortHands(a, b)
	})
	slices.SortFunc(highc, func(a, b string) int {
		return sortHands(a, b)
	})

	twin := 0
	count := 1

	for _, v := range highc {
		twin += count * hands[v]
		count++
	}
	for _, v := range onep {
		twin += count * hands[v]
		count++
	}
	for _, v := range twop {
		twin += count * hands[v]
		count++
	}
	for _, v := range threek {
		twin += count * hands[v]
		count++
	}
	for _, v := range fullh {
		twin += count * hands[v]
		count++
	}
	for _, v := range fourk {
		twin += count * hands[v]
		count++
	}
	for _, v := range fivek {
		twin += count * hands[v]
		count++
	}

	fmt.Println(twin)
}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	hands := map[string]int{}

	fivek := []string{}
	fourk := []string{}
	fullh := []string{}
	threek := []string{}
	twop := []string{}
	onep := []string{}
	highc := []string{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		key := line[0]
		value, _ := strconv.Atoi(line[1])
		hands[key] = value

		hand := strings.Split(key, "")
		mhand := map[string]int{}

		for _, v := range hand {
			mhand[v] += 1
		}

		if len(mhand) == 1 {
			fivek = append(fivek, key)
			continue
		}
		if len(mhand) == 5 {
			if _, ok := mhand["J"]; ok {
				onep = append(onep, key)
			} else {
				highc = append(highc, key)
			}
			continue
		}
		if len(mhand) == 4 {
			if _, ok := mhand["J"]; ok {
				threek = append(threek, key)
			} else {
				onep = append(onep, key)
			}
			continue
		}
		if len(mhand) == 2 {
			if _, ok := mhand["J"]; ok {
				fivek = append(fivek, key)
			} else {
				if mhand[string(key[0])] == 1 || mhand[string(key[0])] == 4 {
					fourk = append(fourk, key)
				} else {
					fullh = append(fullh, key)
				}
			}
			continue
		}
		if len(mhand) == 3 {
			if v, ok := mhand["J"]; ok {
				if v == 3 || v == 2 {
					fourk = append(fourk, key)
				} else {
					if mhand[string(key[0])] == 2 || mhand[string(key[1])] == 2 {
						fullh = append(fullh, key)
					} else {
						fourk = append(fourk, key)
					}
				}
			} else {
				if mhand[string(key[0])] == 2 || mhand[string(key[1])] == 2 {
					twop = append(twop, key)
				} else {
					threek = append(threek, key)
				}
			}
			continue
		}

	}
	//	fmt.Println("highc", highc)
	//	fmt.Println("onep", onep)
	//	fmt.Println("twop", twop)
	//	fmt.Println("threek", threek)
	//	fmt.Println("fullh", fullh)
	//	fmt.Println("fourk", fourk)
	//	fmt.Println("fivek", fivek)

	slices.SortFunc(fivek, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(fourk, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(fullh, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(threek, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(twop, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(onep, func(a, b string) int {
		return sortHands2(a, b)
	})
	slices.SortFunc(highc, func(a, b string) int {
		return sortHands2(a, b)
	})

	twin := 0
	count := 1

	for _, v := range highc {
		twin += count * hands[v]
		count++
	}
	for _, v := range onep {
		twin += count * hands[v]
		count++
	}
	for _, v := range twop {
		twin += count * hands[v]
		count++
	}
	for _, v := range threek {
		twin += count * hands[v]
		count++
	}
	for _, v := range fullh {
		twin += count * hands[v]
		count++
	}
	for _, v := range fourk {
		twin += count * hands[v]
		count++
	}
	for _, v := range fivek {
		twin += count * hands[v]
		count++
	}

	fmt.Println(twin)
}

func sortHands(h1, h2 string) int {
	for i := 0; i < len(h1); i++ {
		if cards[string(h1[i])] < cards[string(h2[i])] {
			return -1
		}
		if cards[string(h1[i])] > cards[string(h2[i])] {
			return 1
		}
	}
	return 0
}

func sortHands2(h1, h2 string) int {
	for i := 0; i < len(h1); i++ {
		if cards2[string(h1[i])] < cards2[string(h2[i])] {
			return -1
		}
		if cards2[string(h1[i])] > cards2[string(h2[i])] {
			return 1
		}
	}
	return 0
}

func main() {
	p2()
}
