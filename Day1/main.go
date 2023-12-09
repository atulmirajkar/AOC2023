package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile()
}

/*
   part 2
   one - nine
   can I just update the input file, replace one by 1, two by 2, three by 3...
one
two
three
four
five
six
seven
eight
nine

eightwo
if I replace two - eigh2. then I cannot convert 8 which we want here

s.Index

can I update the input file, replace 1 by one, 2 by two?


*/

func constructTwoDigit(s string) int {
	first := getFirstDigit(s)
	second := getLastDigit(s)

	var num int
	var err error
	if num, err = strconv.Atoi(first + second); err != nil {
		log.Fatal(err)
	}
	return num
}

func getFirstDigit(s string) string {
	firstNum, firstNumPos := getFirstNumDigit(s)
	firstStr, firstStrPos := getFirstStringDigit(s)

	if firstNumPos == -1 {
		return firstStr
	} else if firstStrPos == -1 {
		return firstNum
	} else if firstNumPos < firstStrPos {
		return firstNum
	}
	return firstStr
}

func getLastDigit(s string) string {
	lastNum, lastNumPos := getLastNumDigit(s)
	lastStr, lastStrPos := getLastStringDigit(s)

	if lastNumPos == -1 {
		return lastStr
	} else if lastStrPos == -1 {
		return lastNum
	} else if lastNumPos > lastStrPos {
		return lastNum
	}
	return lastStr
}

func getFirstStringDigit(s string) (string, int) {
	md := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	minPos := math.MaxInt
	digit := ""
	for k, v := range md {
		pos := strings.Index(s, k)
		if pos == -1 {
			continue
		}
		if pos < minPos {
			minPos = pos
			digit = v
		}
	}

	if minPos == math.MaxInt {
		return "", -1
	}

	return digit, minPos
}

func getLastStringDigit(s string) (string, int) {
	md := map[string]string{"eno": "1", "owt": "2", "eerht": "3", "ruof": "4", "evif": "5", "xis": "6", "neves": "7", "thgie": "8", "enin": "9"}
	minPos := math.MaxInt
	s = reverseString(s)
	digit := ""
	for k, _ := range md {
		pos := strings.Index(s, k)
		if pos == -1 {
			continue
		}
		if pos < minPos {
			minPos = pos
			digit = k
		}
	}

	if minPos == math.MaxInt {
		return "", -1
	}

	return md[digit], len(s) - minPos - len(digit)
}

func reverseString(s string) string {
	ba := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		ba[i], ba[j] = ba[j], ba[i]
	}
	return string(ba)
}

func getFirstNumDigit(s string) (string, int) {
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return string(s[i]), i
		}
	}
	return "", -1
}

func getLastNumDigit(s string) (string, int) {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return string(s[i]), i
		}
	}
	return "", -1
}

func readFile() {
	var file *os.File
	var err error

	if file, err = os.Open("./input.txt"); err != nil {
		log.Fatal(err)
	}

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num := constructTwoDigit(line)
		fmt.Println(line+":", num)
		sum += num
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
