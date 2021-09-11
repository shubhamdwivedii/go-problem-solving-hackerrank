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
 * Complete the 'biggerIsGreater' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING w as parameter.
 */

func biggerIsGreater(w string) string {
	// Write your code here
	charArr := make([]int, len(w))

	for i, ch := range w {
		charArr[i] = int(ch)
	}

	var stack []int
	swap := -1

	for i := len(charArr) - 1; i > 0; i-- {
		if charArr[i] <= charArr[i-1] {
			stack = append(stack, charArr[i])
		} else {
			swap = i - 1
			stack = append(stack, charArr[i])
			break
		}
	}
	if swap == -1 {
		return "no answer"
	}

	charArr = charArr[:swap+1]

	minIdx := -1
	for i, ch := range stack {
		if charArr[swap] < ch {
			minIdx = i
			break
		}
	}

	charArr[swap], stack[minIdx] = stack[minIdx], charArr[swap]

	fmt.Println(charArr, stack)

	for _, ch := range stack {
		charArr = append(charArr, ch)
	}

	output := ""

	for _, ch := range charArr {
		output += string(ch)
	}

	return output
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		w := readLine(reader)

		result := biggerIsGreater(w)

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
