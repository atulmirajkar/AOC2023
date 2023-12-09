package main

import (
	"fmt"
)

/*
input: time and max dist recorded
total time 7 sec
t  rt  d
0   7   0
1   6   6
2   5   10
3   4   12

4   3   12
5   2   10
6   1   6
7   0   0

total time 2
t  rt  d
0   2   0
1   1   1
2   0   0

t * rt = d
t * (tt-t) = d

number of ways to win given max dist recorded
traverse half the total time (7/2 +1)
if time is odd - result is count *2
else result = count * 2 -1




*/

type tuple struct {
	time uint64
	dist uint64
}

func main() {
	// a := []tuple{{7, 9}, {15, 40}, {30, 200}}
	//a := []tuple{{53, 333}, {83, 1635}, {72, 1289}, {88, 1532}}
	a := []tuple{{53837288, 333163512891532}}
	total := 1
	for _, t := range a {
		total = total * getWinningCnt(t)
	}
	fmt.Println(total)
}

func getWinningCnt(t tuple) int {
	cnt := 0
	var i uint64
	for i = 0; i <= t.time/2; i++ {
		if i*(t.time-i) > t.dist {
			cnt++
		}
	}

	if t.time%2 == 0 {
		return (cnt * 2) - 1
	}
	return cnt * 2
}
