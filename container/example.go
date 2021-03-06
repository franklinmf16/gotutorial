package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	fmt.Println(lastOccurred)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}


func main() {
	str := "abccczdabcbb"
	substring := lengthOfLongestSubstring(str)
	fmt.Print(substring)
}
