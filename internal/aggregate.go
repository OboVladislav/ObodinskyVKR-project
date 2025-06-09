package additionalFunctions

import (
	"sync"
)

type NumericResultInt struct {
	Sum   int
	Max   int
	Min   int
	Count int
}

func mapWorker(records []int, out chan<- any, wg *sync.WaitGroup) {
	defer wg.Done()

	first := true
	var (
		sumInt int
		maxInt int
		minInt int
		count  int
	)

	for _, rec := range records {
		v := rec

		if first {
			sumInt, maxInt, minInt = v, v, v
			first = false
		} else {
			if v > maxInt {
				maxInt = v
			}
			if v < minInt {
				minInt = v
			}
			sumInt += v
		}
		count++

	}
	out <- NumericResultInt{Sum: sumInt, Max: maxInt, Min: minInt, Count: count}

}

func reduce(out <-chan any, numChunks int) any {
	var (
		first = true

		totalSumInt int
		totalMaxInt int
		totalMinInt int
		totalCount  int
	)

	for i := 0; i < numChunks; i++ {
		cur := <-out
		switch val := cur.(type) {
		case NumericResultInt:
			if first {
				totalSumInt = val.Sum
				totalMaxInt = val.Max
				totalMinInt = val.Min
				totalCount = val.Count
				first = false
			} else {
				if val.Max > totalMaxInt {
					totalMaxInt = val.Max
				}
				if val.Min < totalMinInt {
					totalMinInt = val.Min
				}
				totalSumInt += val.Sum
				totalCount += val.Count
			}
		}
	}
	avg := float32(totalSumInt) / float32(totalCount)
	return struct {
		Avg   float32
		Max   int
		Min   int
		Count int
	}{
		Avg:   avg,
		Max:   totalMaxInt,
		Min:   totalMinInt,
		Count: totalCount,
	}
}

func ParallelAggregation_MapReduce(records *[]int, workers int) any {
	if len(*records) == 0 {
		return nil
	}

	// if workers <= 0 || workers > runtime.NumCPU() {
	// 	workers = runtime.NumCPU()
	// }

	chunkSize := (len(*records) + workers - 1) / workers
	out := make(chan any, workers)

	var wg sync.WaitGroup
	for i := 0; i < len(*records); i += chunkSize {
		end := i + chunkSize
		if end > len(*records) {
			end = len(*records)
		}
		wg.Add(1)
		go mapWorker((*records)[i:end], out, &wg)
	}

	wg.Wait()
	close(out)

	return reduce(out, workers)
}

func SequentialAggregate(data []int) any {
	first := true

	var (
		sumInt int
		maxInt int
		minInt int
		count  int
	)

	for _, rec := range data {
		v := rec

		if first {
			sumInt, maxInt, minInt = v, v, v

			first = false
		} else {
			if v > maxInt {
				maxInt = v
			}
			if v < minInt {
				minInt = v
			}
			sumInt += v
		}
		count++
	}

	avg := float32(sumInt) / float32(count)
	return struct {
		Avg   float32
		Max   int
		Min   int
		Count int
	}{
		Avg:   avg,
		Max:   maxInt,
		Min:   minInt,
		Count: count,
	}

}
