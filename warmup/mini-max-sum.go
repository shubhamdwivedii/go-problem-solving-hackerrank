package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	sums := []int64{0, 0, 0, 0, 0}
	max := sums[0]
	min := sums[0]

	for i, _ := range sums {
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			} else {
				sums[i] += int64(arr[j])
			}
		}
		if i == 0 {
			min = sums[0]
		}
		if sums[i] > max {
			max = sums[i]
		}
		if sums[i] < min {
			min = sums[i]
		}

	}

	fmt.Println(min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
