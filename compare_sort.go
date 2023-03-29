package main

import (
	"fmt"
	"go-sorting/functions"
	"time"
)

func main() {
	maxNumCount := 1000000
	n := functions.NumGenerator(maxNumCount, maxNumCount*10)

	copy := copySlice(n)
	timeStart := time.Now()
	functions.Insertion(copy)
	finished := time.Now().Sub(timeStart).Seconds()
	fmt.Printf("Insertion finished After %f\n", finished)

	copy = copySlice(n)
	timeStart = time.Now()
	functions.MergeSort(copy)
	finished = time.Now().Sub(timeStart).Seconds()
	fmt.Printf("Merge finished After %f\n", finished)

	copy = copySlice(n)
	timeStart = time.Now()
	functions.CountSort(copy)
	finished = time.Now().Sub(timeStart).Seconds()
	fmt.Printf("Count finished After %f\n", finished)

	copy = copySlice(n)
	timeStart = time.Now()
	functions.Selection(copy)
	finished = time.Now().Sub(timeStart).Seconds()
	fmt.Printf("Selection finished After %f\n", finished)

	copy = copySlice(n)
	timeStart = time.Now()
	functions.Bubble(n)
	finished = time.Now().Sub(timeStart).Seconds()
	fmt.Printf("Bubble finished After %f\n", finished)

}

func copySlice(n []int) []int {
	newArr := make([]int, len(n))
	for i, v := range n {
		newArr[i] = v
	}

	return newArr
}
