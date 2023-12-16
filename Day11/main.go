package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
0123
....0
x...1
....2
..x.3

1,0 -->3,2

2+2=4
number of rows betwee two rows that does not have any point
number of cols between tow cols that do not have any points

rows and cols that do not have any xs are double counted
distance between each pair of x


observation:
can use manhattan distance


*/

func main() {
	sa := readFile()
	var ba [][]byte
	for i := range sa {
		ba = append(ba, []byte(sa[i]))
	}

	// fmt.Println(part1(ba, 2))
	fmt.Println(part1(ba, 1000000))
}

type point struct {
	r int
	c int
}

func part1(ba [][]byte, scale int) int {
	rm := make(map[int]interface{})
	cm := make(map[int]interface{})
	var ta []point
	for i := range ba {
		for j := range ba[i] {
			if string(ba[i][j]) == "#" {
				rm[i] = nil
				cm[j] = nil
				ta = append(ta, point{i, j})
			}
		}
	}

	sr := make([]int, len(ba))
	rowSum := 0

	for i := range ba {
		if _, prs := rm[i]; !prs {
			rowSum++
		}
		sr[i] = rowSum
	}

	sc := make([]int, len(ba[0]))
	colSum := 0
	for j := range ba[0] {
		if _, prs := cm[j]; !prs {
			colSum++
		}
		sc[j] = colSum
	}

	sum := 0
	for i := 0; i < len(ta); i++ {
		for j := i + 1; j < len(ta); j++ {
			sum += dist(ta, sr, sc, i, j, scale)
		}
	}
	return sum
}

/*
........x........x..........x.......

	0123456788
*/
func dist(ta []point, sr, sc []int, i, j int, scale int) int {

	emptyRows := abs(sr[ta[i].r] - sr[ta[j].r])
	emptyCols := abs(sc[ta[i].c] - sc[ta[j].c])
	dist := myAbs(ta[i].r, ta[j].r) + myAbs(ta[i].c, ta[j].c) + ((scale * emptyRows) - emptyRows) + ((scale * emptyCols) - emptyCols)

	return dist
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
func myAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func readFile() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var sa []string
	for scanner.Scan() {
		sa = append(sa, scanner.Text())
	}
	return sa

}
