package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type atob struct {
	source uint
	dest   uint
	length uint
}

type seedRange struct {
	start uint
	end   uint
}

func main() {
	// var sa []uint
	// readLine(&sa)
	var sr []seedRange
	readLineTwo(&sr)
	var a [][]atob
	for i := 0; i <= 6; i++ {
		var ab []atob
		readFile(fmt.Sprintf("%d.txt", i), &ab)
		a = append(a, ab)
	}

	//sort the arrays by source
	for i := 0; i <= 6; i++ {
		slices.SortFunc(a[i], func(a, b atob) int {
			return cmp.Compare[uint](a.source, b.source)
		})
	}

	//loop on the seeds
	// var minVal uint
	// minVal = math.MaxUint
	// for _, s := range sa {
	// 	v := s
	//
	// 	for i := 0; i <= 6; i++ {
	// 		v = getDest(v, a[i])
	// 	}
	// 	fmt.Println(v)
	// 	minVal = getMin(minVal, v)
	// }
	// fmt.Println(minVal)

	var minVal uint
	minVal = math.MaxUint
	for _, r := range sr {
		for s := r.start; s <= r.end; s++ {
			v := s

			for i := 0; i <= 6; i++ {
				v = getDest(v, a[i])
			}
			minVal = getMin(minVal, v)
		}
	}
	fmt.Println(minVal)
}

func getMin(a uint, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func getDest(s uint, a []atob) uint {
	for _, r := range a {
		if r.source > s {
			break
		}
		if s >= r.source && s <= (r.source+r.length-1) {
			return r.dest + (s - r.source)
		}
	}
	return s
}

func readLineTwo(sr *[]seedRange) {
	var err error
	var file *os.File

	if file, err = os.Open("./seeds.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	ta := strings.Split(line, " ")
	var sa []uint
	for _, s := range ta {
		v, _ := strconv.ParseUint(s, 10, 32)
		sa = append(sa, uint(v))
	}

	for i := 0; i < len(sa); i = i + 2 {
		*sr = append(*sr, seedRange{start: sa[i], end: sa[i] + sa[i+1] - 1})
	}
}

func readLine(sa *[]uint) {
	var err error
	var file *os.File

	if file, err = os.Open("./seeds.txt"); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	ta := strings.Split(line, " ")
	for _, s := range ta {
		v, _ := strconv.ParseUint(s, 10, 32)
		*sa = append(*sa, uint(v))
	}
}

func readFile(fn string, ma *[]atob) {
	var err error
	var file *os.File

	if file, err = os.Open("./" + fn); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		ta := strings.Split(l, " ")
		source, _ := strconv.ParseUint(ta[1], 10, 32)
		dest, _ := strconv.ParseUint(ta[0], 10, 32)
		length, _ := strconv.ParseUint(ta[2], 10, 32)
		*ma = append(*ma, atob{source: uint(source), dest: uint(dest), length: uint(length)})
	}

}
