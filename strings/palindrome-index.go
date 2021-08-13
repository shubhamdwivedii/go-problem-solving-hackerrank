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
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindromeIndex(s string) int32 {
	// Write your code here

	if checkPalindrome(s, int32(len(s)/2), len(s)%2 == 0) {
		return int32(-1)
	}

	mid := int32(len(s)-1) / 2
	even := int32(len(s)-1)%2 == 0
	for i, _ := range s {
		ss := s[:i] + s[i+1:]
		if checkPalindrome(ss, mid, even) {
			return int32(i)
		}
	}

	return int32(-1)
}

func checkPalindrome(s string, mid int32, even bool) bool {
	front := s[:mid]
	back := s[mid:]
	if !even {
		back = s[mid+1:]
	}

	for i, ch := range front {
		if ch != rune(back[len(back)-1-i]) {
			return false
		}
	}

	return true
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

		result := palindromeIndex(s)

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
