package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFileTwo()
}

func readFile() {
	var err error
	var file *os.File

	if file, err = os.Open("./input.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += parseLine(line)
	}

	fmt.Println(sum)
}

/*
cleanup file
split on |
split on " "
first half put in map
for each in second half check if it is present in map - double points if found
*/
func parseLine(s string) int {
	na := strings.Split(s, "|")
	//winning numbers
	wn := strings.Split(na[0], "-")
	// my numbers
	mn := strings.Split(na[1], "-")

	winMap := make(map[int]interface{})
	for _, v := range wn {
		val, _ := strconv.Atoi(strings.Trim(v, " "))
		winMap[val] = nil
	}

	sum := 1
	for _, v := range mn {
		val, _ := strconv.Atoi(strings.Trim(v, " "))
		if _, prs := winMap[val]; prs {
			sum = sum * 2
		}
	}

	sum = sum / 2
	return sum
}

/*
visualize part 2


can we do a 2 pass solution

pass 1
solve and store for each card

pass 2
0   1   2   3   4   5
5   1   2   0   1   0


rec(start, val)
    foreach idx = start till start + val -1
        return 1 + rec(idx+1, memo[val])


*/

func readFileTwo() {
	var err error
	var file *os.File

	if file, err = os.Open("./input.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//memo array
	var ma []int
	for scanner.Scan() {
		line := scanner.Text()
		val := evalLineTwo(line)
		ma = append(ma, val)
	}

	memo := append([]int{len(ma)}, ma...)
	fmt.Println(rec(memo, 0))

}

func rec(memo []int, idx int) int {
	//recurse
	val := 0
	for i := idx + 1; i <= idx+memo[idx]; i++ {
		val += rec(memo, i)
	}
	return 1 + val
}

func evalLineTwo(s string) int {
	na := strings.Split(s, "|")
	//winning numbers
	wn := strings.Split(na[0], "-")
	// my numbers
	mn := strings.Split(na[1], "-")

	winMap := make(map[int]interface{})
	for _, v := range wn {
		val, _ := strconv.Atoi(strings.Trim(v, " "))
		winMap[val] = nil
	}

	sum := 0
	for _, v := range mn {
		val, _ := strconv.Atoi(strings.Trim(v, " "))
		if _, prs := winMap[val]; prs {
			sum += 1
		}
	}

	return sum
}
