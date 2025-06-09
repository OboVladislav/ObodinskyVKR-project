package additionalFunctions

import (
	"sort"
	"sync/atomic"
)

const minSize = 10000

var (
	pCounter  int64
	pCounterQ int64
)

func mergePar(part []int) {
	merged := make([]int, 0, len(part))
	i := 0
	j := len(part) / 2
	for i < len(part)/2 && j < len(part) {
		if part[i] <= part[j] {
			merged = append(merged, part[i])
			i++
		} else {
			merged = append(merged, part[j])
			j++
		}
	}

	for i < len(part)/2 {
		merged = append(merged, part[i])
		i++
	}
	for j < len(part) {
		merged = append(merged, part[j])
		j++
	}

	copy(part, merged)
}

func SequentialMergeSort(records []int) {
	if len(records) < 1000 {
		// SequentialMergeSort(records)
		sort.Ints(records)
		return
	}
	half := len(records) / 2

	SequentialMergeSort(records[:half])
	SequentialMergeSort(records[half:])

	mergePar(records)
}

func sortPart(part []int, done chan int, workers int) {
	defer func() {
		done <- 0
	}()

	if len(part) < minSize {
		sort.Ints(part)
		return
	}

	half := len(part) / 2
	left := make(chan int)
	right := make(chan int)

	current := atomic.LoadInt64(&pCounter)
	if current <= int64(workers-2) {
		atomic.AddInt64(&pCounter, 2)
		go sortPart(part[:half], left, workers)
		go sortPart(part[half:], right, workers)
	} else if current <= int64(workers-1) {
		atomic.AddInt64(&pCounter, 1)
		go sortPart(part[:half], left, workers)
		SequentialMergeSort(part[half:])
		<-left
		mergePar(part)
		return
	} else {
		SequentialMergeSort(part)
		return
	}

	<-left
	<-right
	mergePar(part)
}
func ParallelMergeSort(records []int, workers int) {
	atomic.StoreInt64(&pCounter, 0)
	done := make(chan int)
	go sortPart(records, done, workers)
	<-done
}

func partition(records []int, start, end int) int {
	pivot := records[end]
	left := start - 1
	for right := start; right < end; right++ {
		if records[right] <= pivot {
			left++
			records[left], records[right] = records[right], records[left]
		}
	}
	records[left+1], records[end] = records[end], records[left+1]
	return left + 1
}

func SequentialQuicksort(input []int, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}

	if endIndex-startIndex+1 < minSize {
		sort.Ints(input[startIndex : endIndex+1])
		return
	}

	pivotIndex := partition(input, startIndex, endIndex)
	SequentialQuicksort(input, startIndex, pivotIndex-1)
	SequentialQuicksort(input, pivotIndex+1, endIndex)
}

func ParallelQuicksort(records []int, workers int) {
	atomic.StoreInt64(&pCounterQ, 0)
	done := make(chan bool)

	go ParallelQuicksortIter(records, workers, 0, len(records)-1, done)

	<-done
}

func ParallelQuicksortIter(records []int, workers, start, end int, done chan bool) {
	defer func() {
		if done != nil {
			done <- true
		}
	}()

	if start >= end {
		return
	}

	if end-start+1 < minSize {
		sort.Ints(records[start : end+1])
		return
	}

	pivotIndex := partition(records, start, end)

	var leftDone chan bool
	var runParallel bool

	if atomic.AddInt64(&pCounterQ, 1) <= int64(workers) {
		runParallel = true
		leftDone = make(chan bool)
		go func() {
			ParallelQuicksortIter(records, workers, start, pivotIndex-1, leftDone)
			atomic.AddInt64(&pCounterQ, -1)
		}()
	} else {
		atomic.AddInt64(&pCounterQ, -1)
		SequentialQuicksort(records, start, pivotIndex-1)
	}

	ParallelQuicksortIter(records, workers, pivotIndex+1, end, nil)

	if runParallel {
		<-leftDone
	}
}
