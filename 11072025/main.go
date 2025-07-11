package main

import "fmt"

func CheckLongestSubstr(s string) int {
	seen := make(map[byte]int)
	left, maxLength := 0, 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if idx, ok := seen[ch]; ok && idx >= left {
			left = idx + 1
		}

		seen[ch] = right
		if maxLength < right-left+1 {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

func main() {
	s := "abcbbcdab"
	fmt.Println(s)
	res := CheckLongestSubstr(s)
	fmt.Println("result:", res)
}
