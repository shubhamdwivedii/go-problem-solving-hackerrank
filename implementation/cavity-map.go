package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func cavityMap(grid []string) []string {
	// Write your code here

	rows := len(grid)
	cols := len(grid[0])

	updateMap := make(map[int]string)
	for i := 1; i < rows-1; i++ {
		prevRow := strings.Split(grid[i-1], "")
		row := strings.Split(grid[i], "")
		copyRow := make([]string, len(row))
		res := copy(copyRow, row)
		fmt.Println("copy row", res)
		nextRow := strings.Split(grid[i+1], "")

		for j := 1; j < cols-1; j++ {
			cell, er1 := strconv.Atoi(string(row[j]))
			top, er2 := strconv.Atoi(string(prevRow[j]))
			bott, er3 := strconv.Atoi(string(nextRow[j]))
			left, er4 := strconv.Atoi(string(row[j-1]))
			right, er5 := strconv.Atoi(string(row[j+1]))

			if er1 != nil || er2 != nil || er3 != nil || er4 != nil || er5 != nil {
				log.Fatal("Error", er1, er2, er3, er4, er5)
			}

			if cell > top && cell > bott && cell > left && cell > right {
				// row[j] = "X"
				copyRow[j] = "X"
				updateMap[i] = strings.Join(copyRow, "")
			}

		}
	}

	for k, v := range updateMap {
		grid[k] = v
	}

	return grid
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

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

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
