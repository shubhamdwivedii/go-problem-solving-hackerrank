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
 * Complete the 'beautifulBinaryString' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING b as parameter.
 */

func beautifulBinaryString(b string) int32 {
	// Write your code here

	if len(b) < 3 {
		return int32(0)
	}
	curr := 0
	count := int32(0)
	for curr <= len(b)-3 {
		fmt.Println("currnet", curr, string(b[curr:curr+3]))
		if string(b[curr]) == "0" && string(b[curr+1]) == "1" && string(b[curr+2]) == "0" {
			count += 1
			curr += 3
		} else {
			curr++
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

	if n < 0 {
		fmt.Println(n)
	}

	b := readLine(reader)

	result := beautifulBinaryString(b)

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
