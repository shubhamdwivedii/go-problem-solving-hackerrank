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
 * Complete the 'anagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func anagram(s string) int32 {
	// Write your code here

	if len(s)%2 != 0 {
		return int32(-1)
	}

	breakPoint := len(s) / 2

	front := s[:breakPoint]
	back := s[breakPoint:]

	fmt.Println(front, back)

	charMap := make(map[string]int32)

	for _, ch := range front {
		if _, ok := charMap[string(ch)]; ok {
			charMap[string(ch)]++
		} else {
			charMap[string(ch)] = 1
		}
	}

	count := int32(0)

	for _, ch := range back {
		if val, ok := charMap[string(ch)]; ok {
			if val <= 0 {
				count++
			} else {
				charMap[string(ch)]--
			}
		} else {
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

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := anagram(s)

		fmt.Fprintf(writer, "%d\n", result)
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
