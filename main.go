package main

import (
	"fmt"
	"math/rand"
)

func GenerateSequence(dest *[]byte) {
	*dest = make([]byte, 2500)
	for i := range *dest {
		(*dest)[i] = byte(rand.Intn(256))
	}
}

func main() {
	var sequence []byte
	GenerateSequence(&sequence)
	var seriesLengths [][]int = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}
	var seriesLengthsStandard [][]int = [][]int{
		{2267, 1079, 502, 223, 90, 90},
		{2733, 1421, 748, 402, 223, 223},
	}
	onesCount := 0
	prevDigit := byte(0)
	curSerLen := 1
	curDigit := byte(0)
	maxSerLen := 0
	for _, el := range sequence {
		for i := 7; i >= 0; i-- {
			curDigit = (el >> i) % 2
			if prevDigit != curDigit {
				seriesLengths[prevDigit][min(curSerLen-1, 5)]++
				maxSerLen = max(maxSerLen, curSerLen)
				curSerLen = 0
			}
			onesCount += int(curDigit)
			prevDigit = curDigit
			curSerLen++
		}
	}
	maxSerLen = max(maxSerLen, curSerLen)
	fmt.Printf("Ones count: %v\n", onesCount)
	if onesCount > 9653 && onesCount < 10347 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	fmt.Printf("\nLongest series: %v\n", maxSerLen)
	if maxSerLen < 37 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	fmt.Println("\nNumber of series of different lengths:")
	fmt.Println(seriesLengths)
	failFlag := false

	for i := 0; i < 6; i++ {
		if seriesLengths[0][i] < seriesLengthsStandard[0][i] ||
			seriesLengths[1][i] < seriesLengthsStandard[0][i] ||
			seriesLengths[0][i] > seriesLengthsStandard[1][i] ||
			seriesLengths[1][i] > seriesLengthsStandard[1][i] {
			fmt.Println("Fail")
			failFlag = true
			break
		}
	}
	if !failFlag {
		fmt.Println("Pass")
	}

	/*for i := range sequence {
		fmt.Println(sequence[i])
	}*/
}
