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
 * Complete the 'funnyString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func funnyString(s string) string {
	// Write your code here

	diffMap := make(map[int32][2]int32)

	for i := 0; i < len(s)-1; i++ {
		diff := math.Abs(float64(s[i]) - float64(s[i+1]))
		diffMap[int32(i)] = [2]int32{int32(diff), -1}
	}

	idx := int32(0)
	for i := len(s) - 1; i > 0; i-- {
		diff := math.Abs(float64(s[i]) - float64(s[i-1]))

		if val, ok := diffMap[idx]; ok {
			diffMap[idx] = [2]int32{val[0], int32(diff)}
		}

		idx++
	}

	for _, v := range diffMap {
		if v[0] != v[1] {
			return "Not Funny"
		}
	}
	return "Funny"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := funnyString(s)

		fmt.Fprintf(writer, "%s\n", result)
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
