package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var content []byte
	var err error
	content, err = ioutil.ReadFile("./input")
	if err != nil {
		fmt.Println("err", err)
	}

	content_string := string(content)
	num_strings := strings.Split(content_string, "\n")
	var nums []int

	for i := 0; i < len(num_strings); i++ {
		num, _ := strconv.Atoi(num_strings[i])
		nums = append(nums, num)
	}

	var increases int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			increases += 1
		}
	}

	fmt.Println("increases", increases)
}
