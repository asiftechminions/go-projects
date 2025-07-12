package main

import (
	"fmt"
	"sort"
)

func GetMergedIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := mergedIntervals[len(mergedIntervals)-1]
		if intervals[i][0] < last[1] {
			last[1] = max(last[1], intervals[i][1])
		} else {
			mergedIntervals = append(mergedIntervals, intervals[i])
		}
	}

	return mergedIntervals

}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intMerged := GetMergedIntervals(intervals)
	if len(intMerged) > 0 {
		fmt.Println("The merged intervals are:", intMerged)
	} else {
		fmt.Println("Intervals is empty")
	}
}
