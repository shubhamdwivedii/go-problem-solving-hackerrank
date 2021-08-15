package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here

	// length := checkSequence(a, 0, 0, 1, 1, 1)
	// return length

	numMap := make(map[int32]int32)

	for _, num := range a {
		if _, ok := numMap[num]; ok {
			numMap[num]++
		} else {
			numMap[num] = 1
		}
	}

	max := int32(0)

	for num, count := range numMap {
		if cc, ok := numMap[num-1]; ok {
			if cc+count > max {
				max = cc + count
			}
		} else {
			if count > max {
				max = count
			}
		}
	}

	return max

}

func checkSequence(a []int32, s int, i int, j int, currentLength int32, maxLength int32) int32 {
	if i < len(a)-1 && s < len(a)-2 {
		if j < len(a) {
			fmt.Println("checking seq a", a, " >>i", a[i], ">>j", a[j], "curr", currentLength, "max", maxLength)
			if math.Abs(float64(a[j]-a[i])) <= 1 {
				currentLength += 1
				if maxLength < currentLength {
					maxLength = currentLength
				}
				return checkSequence(a, s, j, j+1, currentLength, maxLength)
			} else {
				return checkSequence(a, s, i, j+1, currentLength, maxLength)
			}
		} else {
			if i+1 < len(a) && i > s {
				return checkSequence(a, s, s, i+1, 1, maxLength)
			} else {
				return checkSequence(a, s+1, s+1, s+2, 1, maxLength)
			}
		}
	} else {
		return maxLength
	}

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

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
