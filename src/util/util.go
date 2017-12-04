package util

import (
	"io/ioutil"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFileToString(filename string) string {
	file, err := ioutil.ReadFile(filename)
	check(err)
	str := string(file)
	return str
}

func Abs(number int) int {
	if number < 0{
		return -1*number
	}else{
		return number
	}
}

func ReadFileLines(filename string)[]string {
	fileStr := ReadFileToString(filename)
	lines := strings.Split(fileStr,"\n")
	return lines
}
