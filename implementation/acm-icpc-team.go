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
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

func acmTeam(topic []string) []int32 {
	// Write your code here

	maxCount := int32(0)
	countOfMax := int32(0)
	for i := 0; i < len(topic)-1; i++ {
		for j := i + 1; j < len(topic); j++ {
			first := topic[i]
			second := topic[j]

			count := int32(0)

			for in, ch := range first {
				n1, err1 := strconv.Atoi(string(ch))
				n2, err2 := strconv.Atoi(string(second[in]))

				if err1 != nil || err2 != nil {
					log.Fatal("Error", err1, err2)
				}

				if n1+n2 > 0 {
					count++
				}
			}

			if count > maxCount {
				maxCount = count
				countOfMax = int32(1)
			} else if count == maxCount {
				countOfMax++
			}
		}
	}

	return []int32{maxCount, countOfMax}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	if m < 1 {
		fmt.Println(m)
	}

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
