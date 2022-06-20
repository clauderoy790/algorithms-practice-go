package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := []interval{{8, 9}, {1, 8}, {1, 2}, {2, 9}, {2, 4}, {1, 7}, {8, 10}}

	uniques := getUniqueIntervals(intervals)
	fmt.Println("THE UNIQUES ARE: ", uniques)
}

func getUniqueIntervals(intervals []interval) []interval {
	uniques := make([]interval, 0)
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start || (intervals[i].start == intervals[j].start && intervals[i].end < intervals[j].end)
	})

	var currInterval *interval
	for _, ival := range intervals {
		if currInterval == nil {
			currInterval = &interval{ival.start, ival.end}
			continue
		}

		if ival.start == currInterval.start && ival.end > currInterval.end {
			currInterval.end = ival.end
			continue
		}

		if overlaps(ival, currInterval) {
			continue
		}

		uniques = append(uniques, *currInterval)
		currInterval = &interval{ival.start, ival.end}
	}

	if currInterval != nil {
		uniques = append(uniques, *currInterval)
	}

	return uniques
}

func overlaps(ival interval, currInterval *interval) bool {
	return ival.start >= currInterval.start && ival.end <= currInterval.end
}

type interval struct {
	start, end int
}
