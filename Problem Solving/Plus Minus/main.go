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
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
    // Write your code here
    var totalValues int32 = int32(len(arr))
    var positives, negatives, zeroes int32 = 0, 0, 0
    
    // Count positive, negative and zero values
    for _, v := range arr {
        if v > 0 {
            positives++
        } else if v < 0 {
            negatives++
        } else {
            zeroes++
        }
    }
    
    // Print fractions
    fmt.Printf("%.6f\n", float64(positives)/float64(totalValues))
    fmt.Printf("%.6f\n", float64(negatives)/float64(totalValues))
    fmt.Printf("%.6f\n", float64(zeroes)/float64(totalValues))
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

    plusMinus(arr)
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
