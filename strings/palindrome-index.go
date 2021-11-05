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
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Complete the palindromeIndex function below.
func palindromeIndex(s string) int32 {
	l := len(s) - 1
	for i := 0; i <= l; i, l = i+1, l-1 {
		if s[i] == s[l] {
			continue
		}
		if palindrome(s[i+1 : l+1]) {
			return int32(i)
		}
		if palindrome(s[i:l]) {
			return int32(l)
		}
		return -1
	}
	return -1
}

func checkPalindrome(s string, curr int32, broken bool) (int32, bool) {
	fmt.Println("Checkpalin", s, curr, broken)
	if len(s) <= 1 {
		if !broken {
			return -1, true
		}
		return curr, true
	} else {
		if s[0] == s[len(s)-1] {
			if broken {
				return checkPalindrome(s[1:len(s)-1], curr, broken)
			}
			return checkPalindrome(s[1:len(s)-1], curr+1, broken)
		} else {
			if broken {
				return curr, false
			}
			broken = true

			if s[1] == s[len(s)-1] {
				// continue with left omited
			}

			leftRem, lpalin := checkPalindrome(s[1:], curr, broken)
			rightRem, rpalin := checkPalindrome(s[:len(s)-2], int32(len(s))+curr-1, broken)

			if lpalin {
				return leftRem, true
			}
			return rightRem, rpalin

		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := palindromeIndex(s)

		fmt.Fprintf(writer, "%d\n", result)
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
