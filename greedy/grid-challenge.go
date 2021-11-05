package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	// Write your code here
	for idx, str := range grid {
		chars := strings.Split(str, "")
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		grid[idx] = strings.Join(chars, "")
	}

	// sort.Slice(grid, func(i,j int) bool {
	//     return grid[i][0] < grid[j][0]
	// })

	cols := len(grid[0])

	fmt.Println("grid", grid)

	j := 0

	for j < cols {
		comp := grid[0][j]
		for i := 1; i < len(grid); i++ {
			if grid[i][j] < comp {
				return "NO"
			} else {
				comp = grid[i][j]
			}
		}
		j++
	}

	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var grid []string

		for i := 0; i < int(n); i++ {
			gridItem := readLine(reader)
			grid = append(grid, gridItem)
		}

		result := gridChallenge(grid)

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
