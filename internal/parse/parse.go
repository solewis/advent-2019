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
		i, err := strconv.Atoi(l)
		check(err)
		ints = append(ints, i)
	}
	return ints
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
