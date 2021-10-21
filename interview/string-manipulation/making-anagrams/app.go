package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func removeDiff(longer string, shorter string) int32 {
	counter := int32(0)
	newLonger := ""
	for _, l := range longer {
		found := false
		for _, s := range shorter {
			if l == s {
				found = true
				break
			}
		}
		if found {
			newLonger += string(l)
		} else {
			counter++
		}
	}
	lLen := len(newLonger)
	sLen := len(shorter)
	if counter == 0 {
		return counter
	}
	if lLen == sLen {
		counter += removeDiff(shorter, newLonger)
	} else {
		if lLen > sLen {
			counter += removeDiff(newLonger, shorter)
		} else {
			counter += removeDiff(shorter, newLonger)
		}
	}

	return counter
}
func makeAnagrams(a string, b string) int32 {
	// one of 2 string length == 0
		// return length of another
	// do
	// find longer string
		// delete all chars dont exist in shorter one
		// count ++ for each deletion
	// while 2 string not equal
	aLen := int32(len(a))
	bLen := int32(len(b))
	if  aLen == 0 || bLen == 0 {
		if aLen > bLen {
			return aLen
		} else {
			return bLen
		}
	}
	longerStr := b
	shorterStr := a
	if aLen > bLen {
		longerStr = a
		shorterStr = b
	}
	return removeDiff(longerStr, shorterStr)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	aTmp, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	a := strings.Trim(string(aTmp), "\r\n")
	bTmp, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	b := strings.Trim(string(bTmp), "\r\n")
	res := makeAnagrams(a,b)
	fmt.Println(res)
}
