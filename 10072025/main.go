package main

import "fmt"

func CanFormTarget(target int, input []int) []int {
	targetFormer := make(map[int]int)
	for i, v := range input {
		if j, ok := targetFormer[target-v]; ok {
			return []int{j, i}
		}
		targetFormer[v] = i
	}
	return nil
}

func main() {
	input := []int{2, 7, 11, 15}
	target := 19

	if result := CanFormTarget(target, input); result != nil {
		fmt.Println("match found", result)
	} else {
		fmt.Println("No match found")
	}
}
