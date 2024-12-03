package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    // Check if the input is AM time
    isAM := strings.HasSuffix(s, "AM")
    
    // Sanitize the input to cut the AM or PM indicator
    sanitizedTimeStr := strings.TrimSuffix(strings.TrimSuffix(s, "PM"), "AM") 
    
    // Split the input into parts
    parts := strings.Split(sanitizedTimeStr, ":")
    
    // Extract parts
    hour, minute, second := parts[0], parts[1], parts[2]
    var finalHour int64 = 0
    
    if pHour, err := strconv.ParseInt(hour, 10, 32); err == nil {
        finalHour = pHour + 12
        if !isAM && pHour == 12 {
            finalHour = pHour
        }
        if isAM {
            finalHour = (pHour + 12) % 12
        }
    }
    
    stringHour := fmt.Sprintf("%d", finalHour)
    if finalHour < 10 {
        stringHour = fmt.Sprintf("%02d", finalHour)
    }
    
    result := fmt.Sprintf("%s:%s:%s", stringHour, minute, second)
    
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

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
