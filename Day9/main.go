package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1   3   6   10  15  21  28
2   3   4   5   6   7
1   1   1   1   1
0   0   0   0
*/
func main() {
	la := readFile()
	sum := 0
	for _, v := range la {
		sum += evalLine(v)
	}
	fmt.Println(sum)
}

func evalLine(ln string) int {
	sa := strings.Split(ln, " ")
	var va []int
	for _, s := range sa {
		v, _ := strconv.Atoi(s)
		va = append(va, v)
	}
	va = reverse(va)
	var fa []int
	for !isZeros(va) {
		// fmt.Println(va)
		fa = append(fa, va[len(va)-1])
		var ta []int
		for i := 1; i < len(va); i++ {
			ta = append(ta, va[i]-va[i-1])
		}
		va = ta
	}
	// fmt.Println(fa)
	return doSum(fa)
}

func reverse(a []int) []int {
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	return a
}

func doSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
func isZeros(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func readFile() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var sa []string
	for scanner.Scan() {
		ln := scanner.Text()
		sa = append(sa, ln)
	}
	return sa
}
