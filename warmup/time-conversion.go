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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here

	times := strings.Split(s, ":")

	hour, _ := strconv.ParseInt(times[0], 10, 32)

	// last := s[8]

	aOrP := string(s[8])

	// 00 - 12
	// 13 - 23

	if aOrP == "P" && hour < 12 {
		hour += 12
		if hour >= 24 {
			hour = 0
		}
	} else if aOrP == "A" && hour == 12 {
		hour = 0
	}

	times[0] = strconv.Itoa(int(hour))

	res := strings.Join(times, ":")

	res = strings.Trim(res, "AM")
	res = strings.Trim(res, "PM")

	if hour < 10 {
		res = "0" + res
	}

	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
