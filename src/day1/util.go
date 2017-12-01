package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"strconv"
	"io/ioutil"
)

var print = fmt.Println



func check(e error) {
	if e != nil {
		panic(e)
	}
}



func read_file_to_string(filename string) string{
	file, err := ioutil.ReadFile(filename)
	check(err)
	str := string(file)
	return str
}

func readFileLine(filename string) []int{
	file, err := os.Open(filename)
	defer file.Close()

	check(err)

	result := make([]int, 0)
	reader := bufio.NewReader(file)
	reg,err := regexp.Compile("([0-9]+)")
	var line string
	for {
		line, err = reader.ReadString('\n')

		matches := reg.FindAllString(line,-1)
		fmt.Printf ("%v contains number %v\n",line, matches)
		for _,match := range matches{
			number,_ := strconv.Atoi(match)
			result = append(result,number )
		}
		if err != nil {
			break
		}
	}
	return result
}

