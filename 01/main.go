package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input_one, err := os.ReadFile("/home/vasko/advent_of_code/01/input_01")
	if err != nil {
		panic(err)
	}

	input_two, err := os.ReadFile("/home/vasko/advent_of_code/01/input_02")
	if err != nil {
		panic(err)
	}

	first := strings.Split(string(input_one), "\n")
	second := strings.Split(string(input_two), "\n")

	var f []int
	var s []int

	for i := 0; i < len(first); i++ {
		x, err := strconv.Atoi(first[i])
		if err != nil {
			panic(err)
		}
		f = append(f, x)
	}

	for i := 0; i < len(second); i++ {
		x, err := strconv.Atoi(second[i])
		if err != nil {
			panic(err)
		}
		s = append(s, x)
	}

	sort.Ints(f)
	sort.Ints(s)

	var d []int

	for i := 0; i < len(f); i++ {
		var x = f[i]
		var y = s[i]
		if x > y {
			d = append(d, x-y)
		}
		if y > x {
			d = append(d, y-x)
		}
	}

	var sum int

	for i := 0; i < len(d); i++ {
		sum = sum + d[i]
	}

	println(sum)

	var simscores []int

	for i := 0; i < len(f); i++ {

		var x = f[i]
		var count = 0

		for j := 0; j < len(s); j++ {
			if x == s[j] {
				count++
			}
		}

		var ss = x * count
		simscores = append(simscores, ss)

	}

	var simscoressum int
	for i := 0; i < len(simscores); i++ {
		simscoressum = simscoressum + simscores[i]
	}

	println(simscoressum)
}
