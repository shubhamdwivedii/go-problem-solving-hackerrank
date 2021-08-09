package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'marsExploration' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func marsExploration(s string) int32 {
	// Write your code here
	// total := int32(len(s)/3)
	count := int32(0)

	for i := 0; i < len(s); i += 3 {

		if s[i:i+3] != "SOS" {
			seg := s[i : i+3]
			if string(seg[0]) != "S" {
				count++
			}
			if string(seg[1]) != "O" {
				count++
			}
			if string(seg[2]) != "S" {
				count++
			}
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

	s := readLine(reader)

	result := marsExploration(s)

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
