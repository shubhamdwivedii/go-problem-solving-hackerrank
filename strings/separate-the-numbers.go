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
 * Complete the 'separateNumbers' function below.
 *
 * The function accepts STRING s as parameter.
 */

func separateNumbers(s string) {
	// Write your code here

	digits := int32(1)
	curr := int32(0)

	var ff string

	matching := false

	for curr+(digits)-1 < int32(len(s))-digits {
		first := s[curr : curr+digits]
		second := s[curr+digits : curr+(digits*2)]

		all9 := true
		for _, ch := range first {
			if string(ch) != "9" {
				all9 = false
				break
			}
		}

		if all9 {
			second = s[curr+digits : curr+(digits*2)+1]
			digits++
		}

		fNum, err1 := strconv.Atoi(first)
		sNum, err2 := strconv.Atoi(second)

		if fNum+1 == sNum && err1 == nil && err2 == nil && string(first[0]) != "0" && string(second[0]) != "0" {
			if curr == 0 {
				matching = true
				ff = first
			}
			if !all9 {
				curr += digits
			} else {
				curr += digits - 1
			}
		} else {
			if matching == true {
				matching = false
				curr = 0
			}
			if !all9 {
				digits++
			}
		}
	}

	if matching && curr+digits == int32(len(s)) {
		fmt.Println("YES", ff)
	} else {
		fmt.Println("NO")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		separateNumbers(s)
	}
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
