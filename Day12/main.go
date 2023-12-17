package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "#.#.### 1,1,3"
	sa := strings.Split(s, " ")
	ga := getGroups(sa[1])
	ways := 0
	recurse([]byte(sa[0]), ga, &ways, 0)
	fmt.Println(ways)
}
func recurse(ba []byte, ga []int, ways *int, idx int) {
	//exit condition
	if idx == len(ba) {
		if isValid(ba, ga) {
			(*ways)++
		}
	}

	if ba[idx] != '?' {
		recurse(ba, ga, ways, idx+1)
	}

	//recurse
	ba[idx] = '.'
	recurse(ba, ga, ways, idx+1)

	ba[idx] = '#'
	recurse(ba, ga, ways, idx+1)

	ba[idx] = '?'
}

func isValid(ba []byte, ga []int) bool {
	return false
}

func getGroups(s string) []int {
	sa := strings.Split(s, ",")
	var ia []int
	for i := range ia {
		v, _ := strconv.Atoi(sa[i])
		ia = append(ia, v)
	}
	return ia
}
