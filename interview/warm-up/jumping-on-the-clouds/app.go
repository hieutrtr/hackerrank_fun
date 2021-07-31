package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Printf("number of clouds %d", 10)
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var c []int32

    for i := 0; i < int(n); i++ {
        cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
        checkError(err)
        cItem := int32(cItemTemp)
        c = append(c, cItem)
    }
    result := jumpingOnClouds(c)

    fmt.Printf("number of steps %d", result)
}

func jumpingOnClouds(clouds []int32) int32 {
    var steps int32 = 0
    var maxJump int32 = 2
    var next int32 = 2
    for i := range clouds {
        if next == int32(len(clouds)) {
            next -= 1
        }
        if clouds[next] == 1 {
            next -= 1
        }
        if int32(i) == next {
            steps += 1
            next = int32(i) + maxJump
            if next >= int32(len(clouds)) {
                break
            }
        }
    }
    return steps
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