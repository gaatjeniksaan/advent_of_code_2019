package main

import (
	"fmt"
	"github.com/gaatjeniksaan/advent_of_code_2019/helpers"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	helpers.Check(err)
	fmt.Println("hoi")
}
