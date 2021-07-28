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
 * Complete the 'happyLadybugs' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING b as parameter.
 */

func happyLadybugs(b string) string {
	// Write your code here
	colMap := make(map[string]int32)
	emptyCount := int32(0)

	for _, val := range b {
		if string(val) == "_" {
			emptyCount++
			continue
		}
		if count, ok := colMap[string(val)]; ok {
			colMap[string(val)] = count + 1
		} else {
			colMap[string(val)] = int32(1)
		}
	}

	for _, v := range colMap {
		if v < int32(2) {
			return "NO"
		}
	}

	if emptyCount == int32(len(b)) {
		return "YES"
	}

	if emptyCount == int32(0) {
		// check if all happy
		currentColor := string(b[0])
		currentCount := int32(0)
		for _, v := range b {
			if currentColor == string(v) {
				currentCount++
				continue
			} else {
				if currentCount < int32(2) {
					return "NO"
				}
				currentColor = string(v)
				currentCount = int32(1)
			}
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

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		fmt.Println("n", n)
		b := readLine(reader)

		result := happyLadybugs(b)

		fmt.Fprintf(writer, "%s\n", result)
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
