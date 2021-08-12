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
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
 */

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	count := int32(0)

	firstSymb := int32(rune('!'))
	lastSymb := int32(rune('/'))
	hasSymb := false

	firstNum := int32(rune('0'))
	lastNum := int32(rune('9'))
	hasNum := false

	firstLower := int32(rune('a'))
	lastLower := int32(rune('z'))
	hasLower := false

	firstUpper := int32(rune('A'))
	lastUpper := int32(rune('Z'))
	hasUpper := false

	for _, ch := range password {
		switch c := int32(ch); {
		case c >= firstSymb && c <= lastSymb:
			hasSymb = true

		case c >= firstNum && c <= lastNum:
			hasNum = true

		case c >= firstLower && c <= lastLower:
			hasLower = true

		case c >= firstUpper && c <= lastUpper:
			hasUpper = true
		}
	}

	if !hasSymb {
		count++
	}
	if !hasNum {
		count++
	}
	if !hasLower {
		count++
	}
	if !hasUpper {
		count++
	}

	if int32(len(password))+count < int32(6) {
		count += int32(6) - int32(len(password)) - count
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

	password := readLine(reader)

	answer := minimumNumber(n, password)

	fmt.Fprintf(writer, "%d\n", answer)

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
