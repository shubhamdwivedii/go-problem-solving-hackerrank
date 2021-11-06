package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	// Write your code here
	var (
		max   int
		chars = make(map[rune]int, 0)
	)

	for _, r := range []rune(s) {
		chars[r]++
		if chars[r] > max {
			max = chars[r]
		}
	}

	freqs := make([]int, max+1)
	for _, v := range chars {
		freqs[v]++
	}

	var count int
	var prevFreq int
	for j := 0; j < len(freqs); j++ {
		if freqs[j] > 0 {
			count++
			if count > 2 {
				return "NO"
			}
			if prevFreq != 0 {
				switch {
				case prevFreq == 1 && freqs[prevFreq] == 1:
				case j-prevFreq == 1 && freqs[j] == 1:
				default:
					return "NO"
				}
			}
			prevFreq = j
		}
	}

	return "YES"

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

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
