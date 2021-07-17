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
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here

	goodDayCount := int32(0)

	for ii := i; ii <= j; ii++ {
		str := strconv.Itoa(int(ii))
		rev := ""
		for _, c := range str {
			rev = string(c) + rev
		}

		rn, err := strconv.Atoi(rev)

		if err != nil {
			fmt.Print("Error reversing string")
		} else {
			diff := int32(math.Abs(float64(ii - int32(rn))))

			if diff%k == 0 {
				goodDayCount += 1
			}
		}

	}
	return goodDayCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
