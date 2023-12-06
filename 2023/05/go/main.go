package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func p1() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedNrs := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	soil := [][]int{}
	fert := [][]int{}
	water := [][]int{}
	light := [][]int{}
	temp := [][]int{}
	humid := [][]int{}
	loc := [][]int{}
	step := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			step = ""
			continue
		}

		if strings.Contains(line, "map") {
			step = strings.Fields(strings.Split(line, ":")[0])[0]
			continue
		}

		fields := strings.Fields(line)
		dest, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		r, _ := strconv.Atoi(fields[2])

		if step == "seed-to-soil" {
			temp := []int{dest, source, r}
			soil = append(soil, temp)
		}
		if step == "soil-to-fertilizer" {
			temp := []int{dest, source, r}
			fert = append(fert, temp)
		}
		if step == "fertilizer-to-water" {
			temp := []int{dest, source, r}
			water = append(water, temp)
		}
		if step == "water-to-light" {
			temp := []int{dest, source, r}
			light = append(light, temp)
		}
		if step == "light-to-temperature" {
			t := []int{dest, source, r}
			temp = append(temp, t)
		}
		if step == "temperature-to-humidity" {
			temp := []int{dest, source, r}
			humid = append(humid, temp)
		}
		if step == "humidity-to-location" {
			temp := []int{dest, source, r}
			loc = append(loc, temp)
		}

	}

	locations := []int{}
	for _, seed := range seedNrs {
		value, _ := strconv.Atoi(seed)

		for _, v := range soil {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range fert {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range water {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range light {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range temp {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range humid {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		for _, v := range loc {
			if value >= v[1] && value < v[1]+v[2] {
				value = v[0] + (value - v[1])
				break
			}
		}
		locations = append(locations, value)

	}

	min := 0
	for i, l := range locations {
		if i == 0 {
			min = l
		} else if l < min {
			min = l
		}

	}
	fmt.Println(min)

}

func p2() {
	file, err := os.Open("puzzle")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedNrs := strings.Fields(strings.Split(scanner.Text(), ":")[1])

	soil := [][]int{}
	fert := [][]int{}
	water := [][]int{}
	light := [][]int{}
	temp := [][]int{}
	humid := [][]int{}
	loc := [][]int{}
	step := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			step = ""
			continue
		}

		if strings.Contains(line, "map") {
			step = strings.Fields(strings.Split(line, ":")[0])[0]
			continue
		}

		fields := strings.Fields(line)
		dest, _ := strconv.Atoi(fields[0])
		source, _ := strconv.Atoi(fields[1])
		r, _ := strconv.Atoi(fields[2])

		if step == "seed-to-soil" {
			temp := []int{dest, source, r}
			soil = append(soil, temp)
		}
		if step == "soil-to-fertilizer" {
			temp := []int{dest, source, r}
			fert = append(fert, temp)
		}
		if step == "fertilizer-to-water" {
			temp := []int{dest, source, r}
			water = append(water, temp)
		}
		if step == "water-to-light" {
			temp := []int{dest, source, r}
			light = append(light, temp)
		}
		if step == "light-to-temperature" {
			t := []int{dest, source, r}
			temp = append(temp, t)
		}
		if step == "temperature-to-humidity" {
			temp := []int{dest, source, r}
			humid = append(humid, temp)
		}
		if step == "humidity-to-location" {
			temp := []int{dest, source, r}
			loc = append(loc, temp)
		}

	}

	locations := []int{}
	for i := 0; i < len(seedNrs); i++ {
		count, _ := strconv.Atoi(seedNrs[i+1])
		val, _ := strconv.Atoi(seedNrs[i])

		fmt.Println(val)
		for j := 0; j < count; j++ {
			value := val + j
			for _, v := range soil {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range fert {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range water {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range light {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range temp {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range humid {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			for _, v := range loc {
				if value >= v[1] && value < v[1]+v[2] {
					value = v[0] + (value - v[1])
					break
				}
			}
			locations = append(locations, value)

		}
		i++
	}

	min := 0
	for i, l := range locations {
		if i == 0 {
			min = l
		} else if l < min {
			min = l
		}

	}
	fmt.Println(min)

}

func main() {
	p2()
}
