package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var a [][]byte
	readFile(&a)
	// fmt.Println(evalEngine(a))
	fmt.Println(evalEngineTwo(a))
}

/*
......
.4567.
...x..
.4567.
...x..

if * then

	get adjacent numbers
	if there are two multiply and return
*/
func evalEngineTwo(a [][]byte) int {
	tr := len(a)
	tc := len(a[0])
	sum := 0
	for i := 0; i < tr; i++ {
		for j := 0; j < tc; j++ {
			if string(a[i][j]) == "*" {
				num := calculateNumTwo(a, i, j, tr, tc)

				fmt.Printf("i=%d, j=%d, a[i][j]=%s, num = %d \n", i, j, string(a[i][j]), num)
				sum += num
			}
		}
	}
	return sum

}

/*
if not a gear should return 0
if a gear should return num

......
.4567.
...x..
*/
type tuple struct {
	row int
	col int
}

func calculateNumTwo(a [][]byte, i, j, tr, tc int) int {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	var ta []tuple
	for _, d := range dir {
		nr := i + d[0]
		nc := j + d[1]
		if isOutOfBound(nr, nc, tr, tc) {
			continue
		}

		if unicode.IsDigit(rune(a[nr][nc])) {
			ta = append(ta, tuple{nr, nc})
		}
	}

	if len(ta) == 0 {
		return 0
	}

	isGear, num := evalTuples(a, ta, tr, tc)

	if isGear {
		return num
	}

	return 0

}

func evalTuples(a [][]byte, ta []tuple, tr, tc int) (bool, int) {
	//we need to deduplicate the tuples as well
	// we can maintain a key "x:y" for each tuple. if already evaluated, then skip

	m := make(map[string]int)
	for _, t := range ta {
		r := t.row
		c := t.col

		//go to left to start of the number
		for !isOutOfBound(r, c, tr, tc) && unicode.IsDigit(rune(a[r][c])) {
			c--
		}

		c++

		key := strconv.Itoa(r) + ":" + strconv.Itoa(c)
		if _, prs := m[key]; prs {
			continue
		}
		//now r and c are start of the digit
		//evaluate the number
		num := 0
		for !isOutOfBound(r, c, tr, tc) && unicode.IsDigit(rune(a[r][c])) {
			cn, _ := strconv.Atoi(string(a[r][c]))
			num = (num * 10) + cn
			c++
		}

		//update map
		m[key] = num
	}
	if len(m) != 2 {
		return false, 0
	}

	num := 1
	for _, v := range m {
		num = num * v
	}

	return true, num
}

/*
......
.4567.
...x..


1. calculate number and return jump index. Calculate neighbor indices. If is surrounded then add to sum. Update j to jump index

*/

func evalEngine(a [][]byte) int {
	tr := len(a)
	tc := len(a[0])
	sum := 0
	for i := 0; i < tr; i++ {
		for j := 0; j < tc; {
			if unicode.IsDigit(rune(a[i][j])) {
				num, jump := calculateNum(a, i, j, tr, tc)

				fmt.Printf("i=%d, j=%d, a[i][j]=%s, jump = %d, num = %d \n", i, j, string(a[i][j]), jump, num)
				j += jump
				sum += num

			} else {
				j++
			}
		}
	}
	return sum
}

/*
if not surrounded should return 0,jumpVal
if surrounded should return num, jumpVal

......
.4567.
...x..
*/
func calculateNum(a [][]byte, i, j, tr, tc int) (int, int) {
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

	num := 0
	jump := 0
	isSur := false
	for !isOutOfBound(i, j, tr, tc) && unicode.IsDigit(rune(a[i][j])) {
		cn, _ := strconv.Atoi(string(a[i][j]))
		num = (num * 10) + cn

		if !isSur {
			for _, d := range dir {
				nr := i + d[0]
				nc := j + d[1]
				if isOutOfBound(nr, nc, tr, tc) {
					continue
				}

				if a[nr][nc] == '.' || unicode.IsDigit(rune(a[nr][nc])) {
					continue
				}

				isSur = true
				break
			}
		}
		jump++
		j++
	}

	if isSur {
		return num, jump
	}

	return 0, jump

}

func isOutOfBound(i, j, tr, tc int) bool {
	if i < 0 || j < 0 || i >= tr || j >= tc {
		return true
	}
	return false
}

func parseLine(s string) []byte {

	var result []byte
	for i := 0; i < len(s); i++ {
		result = append(result, s[i])
	}
	return result
}

func readFile(a *[][]byte) {
	var file *os.File
	var err error

	if file, err = os.Open("./input.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		*a = append(*a, parseLine(line))

	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
