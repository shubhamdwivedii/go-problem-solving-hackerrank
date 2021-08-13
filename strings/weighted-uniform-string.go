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
 * Complete the 'weightedUniformStrings' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER_ARRAY queries
 */

func weightedUniformStrings(s string, queries []int32) []string {
	// Write your code here

	current := int32(-1)
	charMap := make(map[int32]bool)
	counter := int32(0)

	for _, ch := range s {
		chValue := int32(ch) - int32(rune('a')-1)
		fmt.Println(string(ch), "val", chValue)

		if int32(ch) == current {
			counter++
			if _, ok := charMap[chValue*counter]; !ok {
				charMap[chValue*counter] = true
			}
		} else {
			counter = 1
			current = int32(ch)
			if _, ok := charMap[chValue]; !ok {
				charMap[chValue] = true
			}
		}
	}

	res := make([]string, len(queries))

	for i, q := range queries {
		if _, ok := charMap[q]; ok {
			res[i] = "Yes"
		} else {
			res[i] = "No"
		}
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
