package additionalFunctions

import (
	"math/rand"
)

func randomInt64InRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func GenerateRecord() int {

	record := randomInt64InRange(0, 1000000000000000000)
	return record
}

func Generator(part []int, waiters chan int) {
	for i, _ := range part {
		part[i] = rand.Int()
	}
	waiters <- 0
}

func ParallelGenerateDataset(totalRecords int, workers int) []int {

	res := make([]int, totalRecords)
	recordsPerWorker := totalRecords / workers
	remainder := totalRecords % workers

	waiters := make(chan int, workers)
	for i := 0; i < workers; i++ {
		count := recordsPerWorker
		if i < remainder {
			count++
		}
		idx := i * count
		go Generator(res[idx:idx+recordsPerWorker], waiters)
	}

	for i := 0; i < workers; i++ {
		<-waiters
	}

	return res
}

func SequentialGenerateDataset(totalRecords int) []int {

	res := make([]int, 0, totalRecords)
	for i := 0; i < totalRecords; i++ {
		rec := GenerateRecord()
		res = append(res, rec)
	}
	return res
}
