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
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

func surfaceArea(A [][]int32) int32 {
	// Write your code here

	surface := int32(0)

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			surface += 2 // top and bottom

			// left adjacent
			if j == 0 {
				surface += A[i][j]
			} else {
				diff := A[i][j] - A[i][j-1]
				if diff > 0 {
					surface += diff
				}
			}

			// top adjacent
			if i == 0 {
				surface += A[i][j]
			} else {
				diff := A[i][j] - A[i-1][j]
				if diff > 0 {
					surface += diff
				}
			}

			// bottom adjacent
			if i == len(A)-1 {
				surface += A[i][j]
			} else {
				diff := A[i][j] - A[i+1][j]
				if diff > 0 {
					surface += diff
				}
			}

			// right adjacent
			if j == len(A[i])-1 {
				surface += A[i][j]
			} else {
				diff := A[i][j] - A[i][j+1]
				if diff > 0 {
					surface += diff
				}
			}
		}
	}

	return surface

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	HTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	H := int32(HTemp)

	WTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	W := int32(WTemp)

	var A [][]int32
	for i := 0; i < int(H); i++ {
		ARowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var ARow []int32
		for _, ARowItem := range ARowTemp {
			AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			ARow = append(ARow, AItem)
		}

		if len(ARow) != int(W) {
			panic("Bad input")
		}

		A = append(A, ARow)
	}

	result := surfaceArea(A)

	fmt.Fprintf(writer, "%d\n", result)

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
