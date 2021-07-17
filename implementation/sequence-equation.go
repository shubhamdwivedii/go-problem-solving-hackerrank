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
 * Complete the 'permutationEquation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY p as parameter.
 */

func permutationEquation(p []int32) []int32 {
	// Write your code here
	copy := make([]int32, len(p))

	for i := 1; i <= len(p); i++ {
		for idx, val := range p {
			if val == int32(i) {
				for in, vv := range p {
					if int32(idx+1) == vv {
						copy[i-1] = int32(in + 1)
					}
				}
			}
		}
	}

	return copy
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

	pTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var p []int32

	for i := 0; i < int(n); i++ {
		pItemTemp, err := strconv.ParseInt(pTemp[i], 10, 64)
		checkError(err)
		pItem := int32(pItemTemp)
		p = append(p, pItem)
	}

	result := permutationEquation(p)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
