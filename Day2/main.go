package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes
var total map[string]int

func main() {
	readFile()
}

func isValid(maps []map[string]int) bool {
	for _, m := range maps {
		if m["red"] > 12 {
			return false
		}
		if m["green"] > 13 {
			return false
		}
		if m["blue"] > 14 {
			return false
		}
	}
	return true
}

func getPower(maps []map[string]int) int {
	//minmap
	mm := make(map[string]int)
	for _, m := range maps {
		mm["red"] = max(mm["red"], m["red"])
		mm["green"] = max(mm["green"], m["green"])
		mm["blue"] = max(mm["blue"], m["blue"])
	}
	return mm["red"] * mm["green"] * mm["blue"]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
4 red, 5 blue, 9 green; 7 green, 7 blue, 3 red
*/
func parseLine(s string) []map[string]int {
	var result []map[string]int
	//split on ;
	draws := strings.Split(s, ";")
	for _, d := range draws {
		colors := strings.Split(d, ",")
		dm := make(map[string]int)
		for _, c := range colors {
			c = strings.Trim(c, " ")
			//count and color
			cc := strings.Split(c, " ")
			var err error
			var count int
			if count, err = strconv.Atoi(cc[0]); err != nil {
				log.Fatal(err)
			}
			dm[cc[1]] = count
		}
		result = append(result, dm)
	}
	return result
}

func readFile() {
	var file *os.File
	var err error

	if file, err = os.Open("./input.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	ln := 0
	sum := 0
	for scanner.Scan() {
		ln++
		line := scanner.Text()
		// if !isValid(parseLine(line)) {
		// 	continue
		// }
		// sum += ln
		sum += getPower(parseLine(line))
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
