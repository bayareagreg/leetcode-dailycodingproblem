package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// this solution assumes intervals are sorted by start
	res := [][]int{{intervals[0][0], intervals[0][1]}}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			// overlaps
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			// does not overlap
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	//	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	/*
			[[1,3],[2,6],[8,10],[15,18]]
		[[1,4],[4,5]]
		[[1,1],[2,2],[0,0],[2,3],[1,3],[3,5],[2,3],[3,5]]
		[[2,3],[2,2],[3,3],[1,3],[5,7],[2,2],[4,6]]
		[[1,3],[2,6],[8,10],[8,9],[9,11],[15,18],[2,4],[16,17]]
		[[1012,1136],[1137,1417],[1015,1020]]
		[[1,3]]
		[[1,4],[2,3]]
	*/
	intervals := [][]int{{1, 1}, {2, 2}, {0, 0}, {2, 3}, {1, 3}, {3, 5}, {2, 3}, {3, 5}}
	fmt.Printf("%v\n", merge(intervals))
}
