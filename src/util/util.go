package util

import (
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	"fmt"
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

func Max(a int, b int) int {
	if a > b{
		return a
	}else{
		return b
	}
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
	fileStr = strings.Replace(fileStr,"\r","",-1)
	lines := strings.Split(fileStr,"\n")

	return lines
}

func ReadNumbersFromFile(filename string)[]int {
	fileStr := ReadFileToString(filename)
	numbers := ExtractAllNumbers(fileStr)
	return numbers
}

func ExtractAllNumbers(fileStr string) []int {
	strNumbers := regexp.MustCompile("-?[0-9]+").FindAllString(fileStr, -1)
	//strNumbers := strings.Split(fileStr," ")
	numbers := make([]int, len(strNumbers))
	for i, strNumber := range strNumbers {
		numbers[i], _ = strconv.Atoi(strNumber)
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

func mod(number int,length int)int{
	if number < 0 {
		return mod(length +number, length)
	}else{
		return number % length
	}
}

func revertSubList(list []int, start int, stop int,length int){
	copyList := make([]int,len(list))
	copy(copyList,list)
	for i:=0; i < length;i++{
		index1:= mod(start+i,len(list))
		index2:= mod(stop-i,len(list))
		list[index1] = copyList[index2]
	}
}

func doTheHash(list []int, inputLengths []int, numberRounds int){
	start := 0
	skipSize :=0
	for i:=0; i < numberRounds; i++{
		for _,length :=range inputLengths{
			startInvert := (start) % len(list)
			endInvert := (start+length-1) % len(list)
			revertSubList(list,startInvert,endInvert,length)
			start += length+skipSize
			skipSize++
		}
	}

}

func IsNumber(s string) bool {
	strNumbers := regexp.MustCompile("[0-9]+").FindAllString(s, -1)
	return len(strNumbers) == 1
}

func convertToDenseHash(list []int) []int{
	denseHash := make([]int,len(list)/16)
	for i:=0; i < len(list);i+=16{
		for j:=0; j <16; j++{
			denseHash[i/16] ^= list[i+j]
		}
	}
	return denseHash
}

func convertToHex(list []int) string{
	hexString :=""
	for _,element := range list{
		hexString += fmt.Sprintf("%.02x", element)
	}
	return hexString
}
func generateList(length int) []int{
	intList := make([]int, length)
	for i :=0; i < length; i++{
		intList[i] = i
	}
	return intList
}

func getHashInput(input string) []int {
	byteInput := []byte(input)
	newInput := append(byteInput, 17, 31, 73, 47, 23)
	intLengths := make([]int,len(newInput))
	for i,element := range newInput{
		intLengths[i] = int(element)
	}
	return intLengths
}

func KnotHash(input string) string{
	hashList := generateList(256)
	doTheHash(hashList, getHashInput(input), 64)
	denseHash := convertToDenseHash(hashList)
	return convertToHex(denseHash)
}


func StringToBin(input string)  string {
	binaryString:=""
	for _, c := range input {
		intHex ,_ :=  strconv.ParseUint(string(c), 16, 32)
		binaryString += fmt.Sprintf("%04b",intHex)

	}
	return binaryString
}

func HammingWeight(input string) int{
	numberOnes := 0
	for _,bit := range input{
		if string(bit) =="1"{
			numberOnes++
		}
	}
	return numberOnes
}

func BinToGrid(input string)  string{
	result := ""
	for _,bit := range input {
		if string(bit) == "1"{
			result += "#"
		}else{
			result += "."
		}
	}
	return result
}

func StringToIntArray(input string) []int{
	result := make([]int,len(input))
	for i,char := range input{
		result[i],_ = strconv.Atoi(string(char))
	}
	return result
}