package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'gameOfThrones' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func gameOfThrones(s string) string {
	// Write your code here
	even := false
	if len(s)%2 == 0 {
		even = true
	}

	charCount := make(map[string]int32)

	for _, ch := range s {
		if _, ok := charCount[string(ch)]; ok {
			charCount[string(ch)]++
		} else {
			charCount[string(ch)] = 1
		}
	}

	if even {
		for _, v := range charCount {
			if v%2 != 0 {
				return "NO"
			}
		}
		return "YES"
	} else {
		oddFound := false
		for _, v := range charCount {
			if v%2 != 0 {
				if oddFound {
					return "NO"
				} else {
					oddFound = true
				}
			}
		}
		return "YES"
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := gameOfThrones(s)

	fmt.Fprintf(writer, "%s\n", result)

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
