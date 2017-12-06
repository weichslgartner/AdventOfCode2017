package util

import (
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
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
	lines := strings.Split(fileStr,"\r\n")

	return lines
}

func ReadNumbersFromFile(filename string)[]int {
	fileStr := ReadFileToString(filename)
	strNumbers := regexp.MustCompile("[0-9]+").FindAllString(fileStr,-1)
	//strNumbers := strings.Split(fileStr," ")
	numbers := make([]int,len(strNumbers))
	for i,strNumber := range strNumbers{
		numbers[i],_= strconv.Atoi(strNumber)
	}
	return numbers
}



func ArgMax(intSlice []int) (int,int){
	max :=0
	argmax := 0
	for i,number:= range intSlice{
		if number > max {
			argmax = i
			max = number
		}
	}
	return argmax, max
}

func IntListToString(intSlice []int) string{
	strList := make([]string,len(intSlice))
	for i, number :=range intSlice{
		strList[i] = strconv.Itoa(number)
	}
	signature := strings.Join(strList,",")
	return signature

}

