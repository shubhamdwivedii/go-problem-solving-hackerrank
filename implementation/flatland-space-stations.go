package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"

	// "sort"
	"strconv"
	"strings"
)

func calculateDist(n int32, start int32, direction int32, ssMap map[int32]int32) {
	dist := int32(1)
	fmt.Println("Start", start)
	for start+(dist*direction) < n && start+(dist*direction) >= 0 {
		// for start dist < n {
		fmt.Println("dist", dist)
		idx := start + (dist * direction)
		if idx >= 0 && idx < n {
			if val, ok := ssMap[idx]; ok {
				if val == 0 {
					dist = 0
					start = idx
				}
				if val > dist {
					ssMap[idx] = dist
				} // if value is 0
			} else {
				ssMap[idx] = dist
			}
		}
		dist++
	}
}

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int32) int32 {

	ssMap := make(map[int32]int32)

	fmt.Println("c", c)
	for _, v := range c {
		ssMap[v] = 0
	}

	maxLength := int32(0)
	currentLength := int32(0)
	leadingLength := int32(-1)
	trailingLength := int32(-1)
	for i := int32(0); i < n; i++ {
		fmt.Println("i", i)
		if maxLength < currentLength {
			maxLength = currentLength
		}
		if val, ok := ssMap[i]; ok {
			if val == 0 {
				if leadingLength == -1 {
					leadingLength = currentLength
				}
				currentLength = 0
			}
		} else {
			currentLength += 1
		}
	}
	trailingLength = currentLength
	maxLength = int32((maxLength + 1) / 2)

	if leadingLength > 0 || trailingLength > 0 {
		greater := int32(math.Max(float64(leadingLength), float64(trailingLength)))
		return int32(math.Max(float64(greater), float64(maxLength)))
	}

	return maxLength

	// calculateDist(n, int32(c[0]), int32(1), ssMap)

	// if (len(c) > 1) {
	//     calculateDist(n, int32(c[len(c) - 1]), int32(-1), ssMap)
	// }
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int32

	for i := 0; i < int(m); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := flatlandSpaceStations(n, c)

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
