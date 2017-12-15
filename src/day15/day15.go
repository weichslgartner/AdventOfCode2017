package main

import (
	"fmt"
	"strconv"
)

type  Generator struct{
	currentValue int
	factor int
	divider int //2147483647
	multiplier int
}

func (g *Generator) getNextValue()int{
	newValue := g.currentValue * g.factor
	newValue %= g.divider
	g.currentValue = newValue
	return newValue
}


func (g *Generator) getNextValue2()int{
	newValue := g.getNextValue()
	for newValue % g.multiplier !=0{
		newValue = g.getNextValue()
	}
	return newValue
}

func truncateLowerBits(bitstring string, numberBits int)string{
	if len(bitstring) < numberBits{
		leadingZeros := ""
		for i:=0; i < numberBits -len(bitstring);i++{
			leadingZeros += "0"
		}
		return leadingZeros + bitstring
	}else{
		return bitstring[len(bitstring)-numberBits:]
	}
}

func compareBitStreams(valueA int, valueB int, numberBits int) bool{
	bitsA := truncateLowerBits(strconv.FormatInt(int64(valueA), 2),16)
	bitsB := truncateLowerBits(strconv.FormatInt(int64(valueB), 2),16)
	match := true

	for i:=0; i<numberBits; i++{
		if bitsA[i] != bitsB[i]{
			match = false
			break
		}
	}

	return match
}


func part(generatorA Generator,generatorB Generator, iterations int) {

	numberMatches := 0
	for i := 0; i < iterations; i++ {
		a, b := generatorA.getNextValue2(), generatorB.getNextValue2()
		//	fmt.Println(a,b)
		if compareBitStreams(a, b, 16) {
			numberMatches++
		}

	}
	fmt.Println(numberMatches)
}


func main() {
	generatorA := Generator{883, 16807, 2147483647, 1}
	generatorB := Generator{879, 48271, 2147483647, 1}
	part(generatorA,generatorB,40e6)

	//generatorA = Generator{65, 16807, 2147483647, 4}
	//generatorB = Generator{8921, 48271, 2147483647, 8}
	generatorA = Generator{883, 16807, 2147483647, 4}
	generatorB = Generator{879, 48271, 2147483647, 8}
	part(generatorA,generatorB,5e6)
}
