package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"cmp"

	"slices"
)

/*
A K Q J T 9 ... 2

hand is a set of 5

*/

type tuple struct {
	hand string
	bid  int
}

func main() {
	ta := readFile()
	slices.SortFunc(ta, tCompare)
	fmt.Println(ta)
	sum := 0
	for idx, t := range ta {
		sum += (idx + 1) * t.bid
	}
	fmt.Println(sum)
}

func tCompare(a, b tuple) int {
	va := getHandValTwo(a.hand)
	vb := getHandValTwo(b.hand)
	if va == vb {
		//if same values then check hand individually
		return compareStrings(a.hand, b.hand)
	}
	return cmp.Compare(va, vb)
}

func compareStrings(a, b string) int {
	m := make(map[string]int)
	m["A"] = 14
	m["K"] = 13
	m["Q"] = 12
	//m["J"] = 11
	m["J"] = 1 //part 2
	m["T"] = 10
	m["9"] = 9
	m["8"] = 8
	m["7"] = 7
	m["6"] = 6
	m["5"] = 5
	m["4"] = 4
	m["3"] = 3
	m["2"] = 2
	if a[0] != b[0] {
		return cmp.Compare[int](m[string(a[0])], m[string(b[0])])
	}
	if a[1] != b[1] {
		return cmp.Compare[int](m[string(a[1])], m[string(b[1])])
	}
	if a[2] != b[2] {
		return cmp.Compare[int](m[string(a[2])], m[string(b[2])])
	}
	if a[3] != b[3] {
		return cmp.Compare[int](m[string(a[3])], m[string(b[3])])
	}
	if a[4] != b[4] {
		return cmp.Compare[int](m[string(a[4])], m[string(b[4])])
	}

	return 0
}

type sizeTuple struct {
	cnt  int
	card string
}

/*
There can by multiple Js

# AJJ23

A 3
2 1
3 1

2 J
1 A
1 2
1 3

maybe consider transition states
all different --> can have atmost 1j --> one pair
one pair

	2 Js --> 3 oak
	1 J --> 3 oak
	        can also be 2 pair, but 2 pair has a lower value

two pair

	jj jj a
	jj 11 j
	jj 11 2
	22 11 j
	4 Js --> 5 oak
	3 Js --> 5 oak
	2 Js --> 4 oak
	1 j --> 3 oak + 1 pair = full house

3 oak

	jjj 1 2
	111 j 2
	3 Js --> 4 oak
	1 j --> 4 oak

full house

	jjj 11
	111 jj
	3 Js - 5 oak
	2 Js - 5 oak

4 oak

	JJJJ 1
	1111 J
	4 Js - 5 oak
	1 J - 5 oak

5 oak

	no transition needed

handVal = 6..1
5 of a kind 7
4 oak 6
full house = 3 of a kind + 1 pair 5
3 oak 4
two pair 3
one pair 2
all different 1

observations:
map[card] = cnt
if size = 5

	all different

if size = 4

	1 1 1 2
	one pair

if size = 3

	e.g. 2 2 1, 3 1 1
	two pair
	3 oak

if size = 2

	e.g. 3 2, 4 1
	full house  (3oak + 1pair)
	4 oak

if size == 1

	5 oak
*/
func getHandValTwo(s string) int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}

	// range on map and create a tuple array
	// sort the tuple array
	var sta []sizeTuple
	for k, v := range m {
		sta = append(sta, sizeTuple{cnt: v, card: k})
	}

	//sort in reverse
	slices.SortFunc(sta, func(a, b sizeTuple) int {
		return cmp.Compare[int](b.cnt, a.cnt)
	})

	if len(sta) == 5 {
		//all different
		if _, prs := m["J"]; !prs {
			return 1
		}
		if m["J"] == 1 {
			// one pair
			return 2
		}
	}

	if len(sta) == 4 {
		// one pair

		//transition
		// JJ 1 2 3
		// 11 J 2 3
		// 2 Js --> 3 oak
		// 1 J --> 3 oak
		//         can also be 2 pair, but 2 pair has a lower value
		if _, prs := m["J"]; !prs {
			// one pair
			return 2
		}
		if m["J"] == 2 || m["J"] == 1 {
			//3 OAK
			return 4
		}
	}

	if len(sta) == 3 {
		// 2 2 1 - 2 pair
		if sta[0].cnt == 2 {
			if _, prs := m["J"]; !prs {
				return 3
			}
			// jj jj a
			// jj 11 2
			// 22 11 j
			// 4 Js --> 5 oak
			// 2 Js --> 4 oak
			// 1 j --> 3 oak + 1 pair = full house
			if m["J"] == 4 {
				//5oak
				return 7
			}

			if m["J"] == 2 {
				//4 oak
				return 6
			}

			if m["J"] == 1 {
				// full house
				return 5
			}
		}

		if sta[0].cnt == 3 {
			// 3 1 1 - 3 oak
			if _, prs := m["J"]; !prs {
				return 4
			}
			// jjj 1 2
			// 111 j 2
			// 3 Js --> 4 oak
			// 1 j --> 4 oak
			if m["J"] == 3 || m["J"] == 1 {
				//4 oak
				return 6
			}

		}
	}

	if len(sta) == 2 {
		if sta[0].cnt == 3 {
			// 3 2 - full house
			if _, prs := m["J"]; !prs {
				return 5
			}
			// jjj 11
			// 111 jj
			// 3 Js - 5 oak
			// 2 Js - 5 oak
			if m["J"] == 3 || m["J"] == 2 {
				//5 oak
				return 7
			}

		}
		if sta[0].cnt == 4 {
			// 4 1 - 4 oak
			if _, prs := m["J"]; !prs {
				return 6
			}
			// JJJJ 1
			// 1111 J
			// 4 Js - 5 oak
			// 1 J - 5 oak
			if m["J"] == 4 || m["J"] == 1 {
				//5 oak
				return 7
			}

		}
	}

	// 5 - 5 oak
	if len(sta) == 1 {
		return 7
	}

	return 0 //should not happen

}

/*
AAA23

A 3
2 1
3 1

3 A
1 2
1 3

handVal = 6..1
5 of a kind 7
4 oak 6
full house = 3 of a kind + 1 pair 5
3 oak 4
two pair 3
one pair 2
all different 1

observations:
map[card] = cnt
if size = 5

	all different

if size = 4

	1 1 1 2
	one pair

if size = 3

	e.g. 2 2 1, 3 1 1
	two pair
	3 oak

if size = 2

	e.g. 3 2, 4 1
	full house  (3oak + 1pair)
	4 oak

if size == 1

	5 oak
*/
func getHandVal(s string) int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}

	// range on map and create a tuple array
	// sort the tuple array
	var sta []sizeTuple
	for k, v := range m {
		sta = append(sta, sizeTuple{cnt: v, card: k})
	}

	//sort in reverse
	slices.SortFunc(sta, func(a, b sizeTuple) int {
		return cmp.Compare[int](b.cnt, a.cnt)
	})

	if len(sta) == 5 {
		//all different
		return 1
	}

	if len(sta) == 4 {
		// one pair
		return 2
	}

	if len(sta) == 3 {
		// 2 2 1 - 2 pair
		if sta[0].cnt == 2 {
			return 3
		}
		// 3 1 1 - 3 oak
		if sta[0].cnt == 3 {
			return 4
		}
	}

	if len(sta) == 2 {
		// 3 2 - full house
		if sta[0].cnt == 3 {
			return 5
		}
		// 4 1 - 4 oak
		if sta[0].cnt == 4 {
			return 6
		}
	}

	// 5 - 5 oak
	if len(sta) == 1 {
		return 7
	}

	return 0 //should not happen
}

func readFile() []tuple {
	var ta []tuple
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		la := strings.Split(l, " ")
		v, _ := strconv.Atoi(la[1])
		ta = append(ta, tuple{hand: la[0], bid: v})
	}
	return ta
}
