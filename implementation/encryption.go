package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	// Write your code here
	sqrt := math.Sqrt(float64(len(s)))
	fmt.Println(sqrt)
	row := int32(math.Floor(sqrt))
	col := int32(math.Ceil(sqrt))
	if int32(len(s)) > row*col {
		row += 1
	}
	words := make([]string, row)
	for i := int32(0); i < row; i++ {
		start := i * col
		end := start + col
		if end > int32(len(s)) {
			end = int32(len(s))
		}
		words[i] = s[start:end]
	}
	encoded := make([]string, col)

	for j := int32(0); j < col; j++ {
		for _, wrd := range words {
			if j > int32(len(wrd)-1) {
				continue
			}
			encoded[j] += string(wrd[j])
		}
	}

	return strings.Join(encoded, " ")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

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
