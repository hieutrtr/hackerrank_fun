package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
const (
	Normal int32 = iota
	Leap
	Trans
	)
func checkTypeOfYear(year int32) int32 {
	if year < 1918 {
		if year % 4 == 0 {
			return Leap
		}
		return Normal
	} else if year > 1918 {
		if year % 400 == 0 ||
			(year % 4 == 0 && year % 100 != 0) {
			return Leap
		}
		return Normal
	}
	return Trans
}
func isElementExist(nums []int32, i int32) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}
func main() {
	// a buffer on memory to keep your input
	// you only need one read
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	yearInput, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearInput)
	yearType := checkTypeOfYear(year)
	febDays := 28
	switch yearType {
		case Normal: febDays = 28
		case Leap: febDays = 29
		case Trans: febDays = 15
	}
	fmt.Println(febDays)
	in31Days := []int32{1,3,5,7,8,10,12}
	in30Days := []int32{4,6,9}
	programerDay := int32(febDays)
	month := int32(0)
	for programerDay < 256 {
		month+=1
		if isElementExist(in31Days, month) {
			programerDay += 31
		} else if isElementExist(in30Days, month) {
			programerDay += 30
		}
	}
	moreThan := programerDay - int32(256)
	day := int32(30)
	if isElementExist(in31Days, month) {
		day = int32(31) - moreThan
	} else if isElementExist(in30Days, month) {
		day = int32(30) - moreThan
	}

	timeMonth := fmt.Sprintf("%d", month)
	if month < int32(10) {
		timeMonth = fmt.Sprintf("0%d", month)
	}
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(programerDay)
	fmt.Printf("%d.%s.%d", day,timeMonth,year)
}