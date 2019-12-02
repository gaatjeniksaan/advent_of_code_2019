package main

import (
	"fmt"
	"github.com/gaatjeniksaan/advent_of_code_2019/helpers/helpers"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	helpers.check(err)
	fmt.Println("hoi")
}
