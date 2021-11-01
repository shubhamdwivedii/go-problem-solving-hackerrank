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
 * Complete the 'maxSubarray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func maxSubarray(arr []int32) []int32 {
	// Write your code here
	// Kadane's Algorithm

	curr := arr[0]

	max := curr

	subS := int32(0)
	maxN := arr[0]

	if arr[0] > 0 {
		subS += arr[0]
	}

	ptr := 1

	for ptr < len(arr) {
		if curr+arr[ptr] > arr[ptr] {
			curr = curr + arr[ptr]
		} else {
			curr = arr[ptr]
		}
		if curr > max {
			max = curr
		}

		if arr[ptr] > 0 {
			subS += arr[ptr]
		}
		if arr[ptr] > maxN {
			maxN = arr[ptr]
		}
		ptr++
	}

	if subS == 0 {
		subS = maxN
	}

	return []int32{max, subS}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var arr []int32

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := maxSubarray(arr)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
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
