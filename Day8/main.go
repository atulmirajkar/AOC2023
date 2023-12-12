package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Neighbor struct {
	left  string
	right string
}

func main() {
	la := readFile()
	da := la[0]
	ma := la[2:]

	m := make(map[string]Neighbor)
	for _, s := range ma {
		evalMap(s, &m)
	}

	// cn := "AAA"
	// i := 0
	// steps := 0
	// for i < len(da) {
	//
	// 	if string(da[i]) == "L" {
	// 		cn = m[cn].left
	// 	} else {
	// 		cn = m[cn].right
	// 	}
	// 	steps++
	//
	// 	if cn == "ZZZ" {
	// 		break
	// 	}
	// 	i++
	// 	i = i % len(da)
	// }
	// fmt.Println(steps)

	sn := getStartingNode(m)

	i := 0
	steps := 0
	for i < len(da) {

		for j := range sn {
			sn[j] = getNextNode(sn[j], string(da[i]), m)
		}
		steps++

		if isEnd(sn) {
			break
		}
		i++
		i = i % len(da)
	}
	fmt.Println(steps)
}

func isEnd(na []string) bool {
	for _, v := range na {
		if string(v[2]) != "Z" {
			return false
		}
	}
	return true
}

func getNextNode(cn string, dir string, m map[string]Neighbor) string {

	if dir == "L" {
		cn = m[cn].left
	} else {
		cn = m[cn].right
	}

	return cn
}

func getStartingNode(m map[string]Neighbor) []string {
	var sn []string
	for k := range m {
		if string(k[2]) == "A" {
			sn = append(sn, k)
		}
	}
	return sn
}

func evalMap(s string, m *map[string]Neighbor) {
	sa := strings.Split(s, "=")
	source := strings.Trim(sa[0], " ")
	n := strings.Trim(sa[1], "() ")
	na := strings.Split(n, ",")
	neigh := Neighbor{left: strings.Trim(na[0], " "), right: strings.Trim(na[1], " ")}
	(*m)[source] = neigh
}

func readFile() []string {
	var la []string
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		la = append(la, l)
	}
	return la
}
