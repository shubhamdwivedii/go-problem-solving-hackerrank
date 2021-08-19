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
 * Complete the 'organizingContainers' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts 2D_INTEGER_ARRAY container as parameter.
 */

func organizingContainers(container [][]int32) string {
	// Write your code here

	ballsCount := make(map[int32]int32)
	contCount := make(map[int32]int32)
	ballsChecked := make(map[int32]bool)

	for ctype, row := range container {
		for btype, count := range row {
			if _, ok := contCount[int32(ctype)]; ok {
				contCount[int32(ctype)] += count
			} else {
				contCount[int32(ctype)] = count
			}

			if _, ok := ballsCount[int32(btype)]; ok {
				ballsCount[int32(btype)] += count
			} else {
				ballsCount[int32(btype)] = count
			}
		}
	}

	for _, ccount := range contCount {
		checked := false
		for btype, bcount := range ballsCount {
			if ccount == bcount {
				if _, ok := ballsChecked[btype]; !ok {
					ballsChecked[btype] = true
					checked = true
					break
				}
			}
		}
		if !checked {
			return "Impossible"
		}
	}

	return "Possible"
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
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var container [][]int32
		for i := 0; i < int(n); i++ {
			containerRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var containerRow []int32
			for _, containerRowItem := range containerRowTemp {
				containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
				checkError(err)
				containerItem := int32(containerItemTemp)
				containerRow = append(containerRow, containerItem)
			}

			if len(containerRow) != int(n) {
				panic("Bad input")
			}

			container = append(container, containerRow)
		}

		result := organizingContainers(container)

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
