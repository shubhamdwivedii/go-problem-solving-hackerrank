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
 * Complete the 'fairRations' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY B as parameter.
 */

func fairRations(B []int32) string {
	// Write your code here
	sum := int32(0)
	for _, v := range B {
		sum += v
	}

	if sum%2 != 0 {
		return "NO"
	}

	rations := int32(0)

	for i := len(B) - 1; i > 0; i-- {
		if B[i]%2 == 0 {
			continue
		}
		B[i]++
		B[i-1]++
		rations += 2
	}

	return strconv.Itoa(int(rations))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)

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
