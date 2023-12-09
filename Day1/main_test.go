package main

import "testing"

func TestGetFirstDigit(t *testing.T) {
	var tests = []struct {
		a    string
		want string
	}{
		{"sdaf8dasf9", "8"},
		{"sdaf5sdf7sdaf", "5"},
		{"adsf7dsf8", "7"},
		{"kdkjqdkvgs2", "2"},
	}
	for _, tt := range tests {
		ans := getFirstDigit(tt.a)
		if ans != tt.want {
			t.Errorf("getFirstDigit(%s) = %s; want %s", tt.a, ans, tt.want)
		}
	}
}
func TestGetFirstNumDigit(t *testing.T) {
	var tests = []struct {
		a       string
		want    string
		wantPos int
	}{
		{"sdaf8dasf9", "8", 4},
		{"sdaf5sdf7sdaf", "5", 5},
		{"7adsf7dsf8", "7", 0},
		{"kdkjqdkvgs2", "2", 10},
	}
	for _, tt := range tests {
		ans, pos := getFirstNumDigit(tt.a)
		if ans != tt.want {
			t.Errorf("getFirstNumDigit(%s) = %s %d; want %s %d", tt.a, ans, pos, tt.want, tt.wantPos)
		}
	}
}

func TestGetLastNumDigit(t *testing.T) {
	var tests = []struct {
		a       string
		want    string
		wantPos int
	}{
		{"sdaf8dasf9", "9", 9},
		{"sdaf5sdf7sdaf", "7", 8},
		{"7adsf7dsf8", "8", 9},
		{"kdkjqdkvgs2", "2", 10},
	}
	for _, tt := range tests {
		ans, pos := getLastNumDigit(tt.a)
		if ans != tt.want {
			t.Errorf("getLastNumDigit(%s) = %s %d; want %s %d", tt.a, ans, pos, tt.want, tt.wantPos)
		}
	}
}

func TestGetFirstStringDigit(t *testing.T) {
	var tests = []struct {
		a       string
		want    string
		wantPos int
	}{
		{"eightwo", "8", 0},
		{"two1nine", "2", 0},
		{"eightwothree", "8", 0},
		{"zoneight234", "1", 1},
		{"kdkjqdkvgs2", "", -1},
	}
	for _, tt := range tests {
		ans, pos := getFirstStringDigit(tt.a)
		if ans != tt.want || pos != tt.wantPos {
			t.Errorf("getFirstStringDigit(%s) = %s, %d; want %s, %d", tt.a, ans, pos, tt.want, tt.wantPos)
		}
	}
}
func TestGetLastStringDigit(t *testing.T) {
	var tests = []struct {
		a       string
		want    string
		wantPos int
	}{
		{"eightwo", "2", 4},
		{"two1nine", "9", 4},
		{"eightwothree", "3", 7},
		{"zoneight234", "8", 3},
		{"kdkjqdkvgs2", "", -1},
	}
	for _, tt := range tests {
		ans, pos := getLastStringDigit(tt.a)
		if ans != tt.want || pos != tt.wantPos {
			t.Errorf("getLastStringDigit(%s) = %s, %d; want %s, %d", tt.a, ans, pos, tt.want, tt.wantPos)
		}
	}
}
