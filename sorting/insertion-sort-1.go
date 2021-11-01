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
 * Complete the 'insertionSort1' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY arr
 */

func insertionSort1(n int32, arr []int32) {
    // Write your code here
    for i := len(arr)-1; i > 0; i-- {
        temp := arr[i]
        for j := i-1; j >=0; j-- {
            if temp < arr[j] {
                arr[j+1] = arr[j]
                printArr(arr)
            } else {
                if j+1 != i {
                    arr[j+1] = temp 
                    temp = -1
                    printArr(arr)
                }
                break 
            }
        }
        if temp != -1 && temp < arr[0] {
            arr[0] = temp 
            printArr(arr)
        }
    }
}

func printArr(arr []int32) {
    for _,n := range arr {
        fmt.Printf("%v ", n )
    }
    fmt.Printf("\n")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    insertionSort1(n, arr)
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
