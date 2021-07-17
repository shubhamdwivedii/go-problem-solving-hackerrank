package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

func formingMagicSquare(s [][]int32) int32 {
	// Write your code here

	costs := make([]int32, 8)

	costs[0] = findSequenceCost(s, 4, 6, 2, 8)
	costs[1] = findSequenceCost(s, 6, 4, 2, 8)
	costs[2] = findSequenceCost(s, 4, 6, 8, 2)
	costs[3] = findSequenceCost(s, 6, 4, 8, 2)
	costs[4] = findSequenceCost(s, 2, 8, 4, 6)
	costs[5] = findSequenceCost(s, 2, 8, 6, 4)
	costs[6] = findSequenceCost(s, 8, 2, 4, 6)
	costs[7] = findSequenceCost(s, 8, 2, 6, 4)

	minCost := costs[0]

	for i := 1; i < len(costs); i++ {
		minCost = int32(math.Min(float64(minCost), float64(costs[i])))
	}

	return minCost
}

func findSequenceCost(s [][]int32, d1t int32, d1b int32, d2t int32, d2b int32) int32 { // eg. d1t 4, d1b 6, d2t 2, t2b 8
	costs := [3][3]int32{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	if s[1][1] != 5 {
		costs[1][1] = int32(math.Abs(float64(s[1][1] - 5)))
	}

	costs[0][0] = d1t - s[0][0]
	costs[2][2] = d1b - s[2][2]
	costs[0][2] = d2t - s[0][2]
	costs[2][0] = d2b - s[2][0]

	costs[0][1] = int32(15 - (s[0][0] + costs[0][0] + s[0][2] + costs[0][2] + s[0][1]))
	costs[1][0] = int32(15 - (s[0][0] + costs[0][0] + s[2][0] + costs[2][0] + s[1][0]))
	costs[1][2] = int32(15 - (s[0][2] + costs[0][2] + s[2][2] + costs[2][2] + s[1][2]))
	costs[2][1] = int32(15 - (s[2][0] + costs[2][0] + s[2][2] + costs[2][2] + s[2][1]))

	totalCost := int32(0)
	for _, row := range costs {
		for _, cost := range row {
			totalCost += absolutify(cost)
		}
	}
	return totalCost
}

func absolutify(a int32) int32 {
	return int32(math.Abs(float64(a)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
