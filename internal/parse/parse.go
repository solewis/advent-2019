package parse

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Lines(filename string) []string {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return strings.Split(string(dat), "\n")
}

func Ints(filename, separator string) []int {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var ints []int
	for _, l := range strings.Split(string(dat), separator) {
		ints = append(ints, Int(l))
	}
	return ints
}

func Int(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
