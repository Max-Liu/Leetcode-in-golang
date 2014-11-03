package main

import (
	"fmt"
	"strings"
)

/*
https://oj.leetcode.com/problems/reverse-words-in-a-string/
Given an input string, reverse the string word by word.

For example,
Given s = "the sky is blue",
return "blue is sky the".
*/

func main() {
	str := "the sky is blue 111"
	fmt.Println(reverseWordsInAString(str))
}

func reverseWordsInAString(str string) (finalStr string) {

	strSlice := strings.Split(str, " ")
	strLen := len(strSlice)

	var halfStrSlice []string

	if strLen%2 == 0 {
		halfStrSlice = strSlice[:strLen/2]
	} else {
		halfStrSlice = strSlice[:(strLen-1)/2]
	}

	for key, value := range halfStrSlice {
		temp := value
		strSlice[key] = strSlice[strLen-key-1]
		finalStr = finalStr + strSlice[strLen-key-1] + " "
		strSlice[strLen-key-1] = temp
	}

	for _, value := range strSlice[len(halfStrSlice):] {
		finalStr += value + " "
	}

	return finalStr
}
