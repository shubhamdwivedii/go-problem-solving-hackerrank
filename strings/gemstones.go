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
 * Complete the 'gemstones' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING_ARRAY arr as parameter.
 */

func gemstones(arr []string) int32 {
	// Write your code here

	gemMap := make(map[string]int)

	for _, str := range arr {
		encMap := make(map[string]bool)
		for _, ch := range str {
			if _, ok := gemMap[string(ch)]; ok {
				if _, ok := encMap[string(ch)]; !ok {
					gemMap[string(ch)]++
					encMap[string(ch)] = true
				}
			} else {
				gemMap[string(ch)] = 1
				encMap[string(ch)] = true
			}
		}
	}

	count := int32(0)

	for _, v := range gemMap {
		if v == len(arr) {
			count++
		}
	}

	return count

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []string

	for i := 0; i < int(n); i++ {
		arrItem := readLine(reader)
		arr = append(arr, arrItem)
	}

	result := gemstones(arr)

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
