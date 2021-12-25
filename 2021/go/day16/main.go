package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

var hexbin = map[string]string{"0": "0000", "1": "0001", "2": "0010", "3": "0011", "4": "0100", "5": "0101",
	"6": "0110", "7": "0111", "8": "1000", "9": "1001", "A": "1010", "B": "1011", "C": "1100", "D": "1101", "E": "1110", "F": "1111"}

var bs string

func main() {
	var b bytes.Buffer

	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")

		for _, value := range data {
			b.WriteString(hexbin[value])
		}

	}

	bs = b.String()

	log.Println(bs)
	// log.Println(len(bs))

	first()
}

func first() {
	version, err := strconv.ParseInt(bs[:3], 2, 64)
	if err != nil {
		log.Println(err)
		return
	}

	id, err := strconv.ParseInt(bs[3:6], 2, 64)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(bs[6:7])

	log.Println(version)
	log.Println(id)

}
