package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
part 1
|
-
F7
LJ
.
S

# Every pipe in the main loop connects to its 2 neighbors, including S

# Steps to the farthest position

Observations:

can there be multiple loops? Dont think so

Do DFS.
maintain length
Result should be total length / 2

find neighbors of S - there should be only 2 connnected pipes for pipe
on the left is there a L or a F
on the top is there an F or a 7
..

also need a visited set
*/

type tuple struct {
	x int
	y int
}

func main() {
	var a [][]byte
	sa := readFile()

	sp := tuple{-1, -1}
	for i := 0; i < len(sa); i++ {
		var b []byte
		fmt.Println(sa[i])
		for j := 0; j < len(sa[i]); j++ {

			b = append(b, sa[i][j])
			if string(sa[i][j]) == "S" {
				sp = tuple{i, j}
			}
		}
		a = append(a, b)
	}

	visited := make(map[tuple]int)
	length := 0
	DFS(a, sp, visited, &length, tuple{-1, -1})
	fmt.Println(length)

}

/*
-1 -1 -1 -1
a--b--c--d
|     |
g--f--e

return true if found a cycle
*/

func DFS(a [][]byte, cp tuple, visited map[tuple]int, length *int, pp tuple) bool {
	fmt.Println(pp, cp, string(a[cp.x][cp.y]), *length)
	//exit condition - cycle found
	if visited[cp] == -1 {
		return true
	}

	if visited[cp] == 1 {
		return false
	}

	visited[cp] = -1
	(*length)++
	dir := getDirs(a, cp)
	for i := range dir {
		np := tuple{cp.x + dir[i].x, cp.y + dir[i].y}
		//check if this is the parent pointer - dont traverse back to where we came from
		if np == pp {
			continue
		}

		//skip if out of bound
		if np.x < 0 || np.x > len(a)-1 || np.y < 0 || np.y > len(a[0])-1 {
			continue
		}

		// F-7
		// | |
		// L-J
		cv := string(a[np.x][np.y])
		if cv == "S" {
			return true
		}

		if dir[i] == (tuple{-1, 0}) && cv != "F" && cv != "7" && cv != "|" {
			continue
		} else if dir[i] == (tuple{1, 0}) && cv != "L" && cv != "J" && cv != "|" {
			continue
		} else if dir[i] == (tuple{0, -1}) && cv != "F" && cv != "L" && cv != "-" {
			continue
		} else if dir[i] == (tuple{0, 1}) && cv != "7" && cv != "J" && cv != "-" {
			continue
		}

		//if cycle found return true
		if DFS(a, np, visited, length, cp) {
			return true
		}

	}
	visited[cp] = 1
	(*length)--
	return false

}

// F-7
// | |
// L-J
func getDirs(a [][]byte, cp tuple) []tuple {
	cv := string(a[cp.x][cp.y])
	var dirs []tuple

	if cv == "F" {
		dirs = append(dirs, tuple{0, 1})
		dirs = append(dirs, tuple{1, 0})
	} else if cv == "7" {
		dirs = append(dirs, tuple{0, -1})
		dirs = append(dirs, tuple{1, 0})
	} else if cv == "J" {
		dirs = append(dirs, tuple{0, -1})
		dirs = append(dirs, tuple{-1, 0})
	} else if cv == "L" {
		dirs = append(dirs, tuple{0, 1})
		dirs = append(dirs, tuple{-1, 0})
	} else if cv == "-" {
		dirs = append(dirs, tuple{0, 1})
		dirs = append(dirs, tuple{0, -1})
	} else if cv == "|" {
		dirs = append(dirs, tuple{-1, 0})
		dirs = append(dirs, tuple{1, 0})
	} else if cv == "S" {
		dirs = append(dirs, tuple{-1, 0})
		dirs = append(dirs, tuple{1, 0})
		dirs = append(dirs, tuple{0, 1})
		dirs = append(dirs, tuple{0, -1})
	}

	return dirs
}

func readFile() []string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	var res []string
	for scanner.Scan() {
		ln := scanner.Text()
		res = append(res, ln)
	}
	return res
}
