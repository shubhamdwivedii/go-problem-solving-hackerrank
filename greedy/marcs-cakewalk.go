package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'marcsCakewalk' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY calorie as parameter.
 */

func marcsCakewalk(calorie []int32) int64 {
	// Write your code here
	sort.Slice(calorie, func(i, j int) bool {
		return calorie[i] > calorie[j]
	})

	cake := int64(0)

	for i, cl := range calorie {
		cake += int64(math.Pow(2, float64(i))) * int64(cl)
	}
	return cake
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

	calorieTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var calorie []int32

	for i := 0; i < int(n); i++ {
		calorieItemTemp, err := strconv.ParseInt(calorieTemp[i], 10, 64)
		checkError(err)
		calorieItem := int32(calorieItemTemp)
		calorie = append(calorie, calorieItem)
	}

	result := marcsCakewalk(calorie)

	fmt.Fprintf(writer, "%d\n", result)

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
