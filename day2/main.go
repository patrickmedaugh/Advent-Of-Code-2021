package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input := ReadInput()
	fmt.Println(input[0])
}

func ReadInput() []string {
	var content []byte
	var err error
	content, err = ioutil.ReadFile("./input")
	if err != nil {
		fmt.Println("err", err)
	}

	content_string := string(content)
	return strings.Split(content_string, "\n")
}

func CalcPositions(inputs []string) int {
	horizontal := 0
	depth := 0
	for _, input := range inputs {
		horizontal, depth = CalcPosition(input, horizontal, depth)
	}
	return horizontal * depth
}

func CalcPosition(input string, horizontal int, depth int) (int, int) {
	navigation := strings.Split(input, " ")
	direction := navigation[0]
	weight := navigation[1]

	return horizontal, depth
}
