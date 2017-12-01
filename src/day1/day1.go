package main

import (
	"fmt"

	"strconv"
)



func sum_identical_digits(file_str string) int {
	sum := 0
	var old_char byte = file_str[len(file_str)-1]
	for _, char := range file_str {
		if old_char == byte(char) {

			number, _ := strconv.Atoi(string(char))
			sum += number
		}

		old_char = byte(char)
	}
	return sum
}

func half_way(file_str string) int {
	sum := 0
	length := (len(file_str))
	for index, char := range file_str {
		if file_str[(index+length/2)%length] == byte(char) {

			number, _ := strconv.Atoi(string(char))
			sum += number
		}


	}
	return sum
}

func main() {
	file_str := read_file_to_string("inputs/day1.txt")

	sum := sum_identical_digits(file_str)
	fmt.Println(sum)

	sum = half_way(file_str)
	fmt.Println(sum)
}