package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'fibonacciModified' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER t1
 *  2. INTEGER t2
 *  3. INTEGER n
 */

func fibonacciModified(t1 int32, t2 int32, n int32) *big.Int {
	// Write your code here
	tab := make([]*big.Int, 0)
	tab = append(tab, big.NewInt(int64(t1)))
	tab = append(tab, big.NewInt(int64(t2)))
	for i := 2; i < int(n); i++ {
		b := big.NewInt(1).Mul(tab[i-1], tab[i-1])
		z := big.NewInt(1).Add(tab[i-2], b)
		tab = append(tab, z)
	}
	return tab[n-1]

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	t1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	t1 := int32(t1Temp)

	t2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	t2 := int32(t2Temp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := fibonacciModified(t1, t2, n)

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
