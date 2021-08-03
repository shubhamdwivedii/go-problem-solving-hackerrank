package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'superReducedString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func superReducedString(s string) string {
	// Write your code here

	reduced, furtherReducible := reduce(s)
	for furtherReducible {
		reduced, furtherReducible = reduce(reduced)
	}
	if len(reduced) == 0 {
		return "Empty String"
	}
	return reduced
}

func reduce(s string) (string, bool) {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			a := ""
			if i+1 > 1 {
				a = s[:i]
			}
			b := ""
			if i+1 < len(s)-1 {
				b = s[i+2:]
			}
			return a + b, true
		}
	}
	return s, false
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

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
