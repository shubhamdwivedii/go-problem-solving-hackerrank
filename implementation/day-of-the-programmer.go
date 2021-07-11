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
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here

	days := int32(0)
	mnth := int32(1)
	for mnth <= 12 {
		days += getDaysInMonth(year, mnth)
		if days >= 256 {
			break
		}
		// At this point its the last day of mnth
		mnth++
	}

	dd := ""
	mm := strconv.Itoa(int(mnth))
	yy := strconv.Itoa(int(year))
	date := int32(0)
	if days == 256 {
		date = getDaysInMonth(year, mnth)
	} else {
		diff := days - 256
		date = getDaysInMonth(year, mnth) - diff
	}

	dd = strconv.Itoa(int(date))
	if mnth < 10 {
		mm = "0" + mm
	}
	if date < 10 {
		dd = "0" + dd
	}

	return dd + "." + mm + "." + yy
}

func getDaysInMonth(year int32, month int32) int32 {
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		return 31
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else { // month is 2 febuary
		if year < 1918 {
			if year%4 == 0 {
				return 29
			} else {
				return 28
			}
		} else if year > 1918 {
			if year%400 == 0 {
				return 29
			} else if year%4 == 0 && year%100 != 0 {
				return 29
			} else {
				return 28
			}
		} else { // year is 1918
			return 15
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
