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

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	t := int32(tTemp)

	for i := int32(0); i < t; i++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)
		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		var q []int32

		for i := int32(0); i < n; i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minBribes(n, q)
	}
}

func minBribes(n int32, q []int32) {
	count := int32(0)
	for i := int32(0); i < n; i++ {
		if (q[i] - 1) - i > 2 {
			fmt.Println("Too chaotic")
			return
		}
		for j := int32(math.Max(0, float64(q[i]-int32(2)))); j < i; j++ {
			if q[j] > q[i] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}