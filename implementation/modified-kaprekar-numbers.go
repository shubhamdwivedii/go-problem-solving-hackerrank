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
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

func kaprekarNumbers(p int32, q int32) {
	// Write your code here
	found := false
	for n := p; n <= q; n++ {
		sqr := int64(n) * int64(n)
		str := strconv.FormatInt(sqr, 10)
		sum := int64(0)
		d := int64(0)
		nCopy := n
		for nCopy > 0 {
			nCopy /= 10
			d++
		}

		left, _ := strconv.Atoi(str[:len(str)-int(d)])
		right, err := strconv.Atoi(str[len(str)-int(d):])

		sum = int64(left) + int64(right)

		if right == 0 && err == nil {
			continue
		}

		// fmt.Println("n", n, "sqr", sqr, "left", left, "right", right, "sum", sum)

		if sum == int64(n) {
			if !found {
				found = true
			}
			fmt.Print(n)
			fmt.Print(" ")
		}
	}
	if !found {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
