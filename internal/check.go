package additionalFunctions

import (
	"sort"
	"strconv"
)

func CheckGenerate() (bool, string) {
	sizes := []int{500, 1000, 5000, 10000, 20000, 100000}
	workers := []int{2, 4, 10, 12, 25, 50}

	for i := 0; i < len(sizes); i++ {

		seqData := SequentialGenerateDataset(sizes[i])

		if len(seqData) != sizes[i] {
			return false, "SeqGenNotLens" + strconv.Itoa(sizes[i])
		}
		for _, v := range seqData {
			if v < 0 || v > 1000000000000000000 {
				return false, "SeqGenValNotRange" + strconv.Itoa(sizes[i])
			}
		}

		for j := 0; j < len(workers); j++ {
			ParData := ParallelGenerateDataset(sizes[i], workers[j])
			if len(ParData) != sizes[i] {
				return false, "ParGenNotLens" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			}
			for _, v := range ParData {
				if v < 0 || v > 1000000000000000000 {
					return false, "ParGenValNotRange" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
				}
			}

		}
	}
	return true, ""

}

func CheckFilter() (bool, string) {
	sizes := []int{500, 1000, 5000, 10000, 20000, 100000}
	workers := []int{2, 4, 10, 12, 25, 50}

	for i := 0; i < len(sizes); i++ {
		data := SequentialGenerateDataset(sizes[i])
		seqData := SequentialFilter(data, 15000)

		sort.Ints(seqData)

		for j := 0; j < len(workers); j++ {
			ParData := ParallelFilter(data, workers[j], 15000)

			sort.Ints(ParData)

			if len(seqData) != len(ParData) {
				return false, "DiffLens" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			} else {
				for i, v := range ParData {
					if v != seqData[i] {
						return false, "DiffResultFilter" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
					}
				}
			}

		}
	}
	return true, ""

}

func CheckMergeSort() (bool, string) {
	sizes := []int{500, 1000, 5000, 10000, 20000, 100000}
	workers := []int{2, 4, 10, 12, 25, 50}

	for i := 0; i < len(sizes); i++ {
		dataSeq := SequentialGenerateDataset(sizes[i])
		dataIter := make([]int, sizes[i])
		copy(dataIter, dataSeq)
		SequentialMergeSort(dataSeq)

		for j := 0; j < len(workers); j++ {
			dataPar := make([]int, sizes[i])
			copy(dataPar, dataIter)
			ParallelMergeSort(dataPar, workers[j])

			if len(dataSeq) != len(dataPar) {
				return false, "DiffLens" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			} else {
				for i, v := range dataPar {
					if v != dataSeq[i] {
						return false, "DiffResultSort" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
					}
				}
			}

		}
	}
	return true, ""

}

func CheckQuickSort() (bool, string) {
	sizes := []int{500, 1000, 5000, 10000, 20000, 100000}
	workers := []int{2, 4, 10, 12, 25, 50}

	for i := 0; i < len(sizes); i++ {
		dataSeq := SequentialGenerateDataset(sizes[i])
		dataIter := make([]int, sizes[i])
		copy(dataIter, dataSeq)
		SequentialQuicksort(dataSeq, 0, len(dataSeq)-1)

		for j := 0; j < len(workers); j++ {
			dataPar := make([]int, sizes[i])
			copy(dataPar, dataIter)
			ParallelQuicksort(dataPar, workers[j])

			if len(dataSeq) != len(dataPar) {
				return false, "DiffLens" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			} else {
				for i, v := range dataPar {
					if v != dataSeq[i] {
						return false, "DiffResultSort" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
					}
				}
			}

		}
	}
	return true, ""

}

func CheckAggregate() (bool, string) {
	sizes := []int{500, 1000, 5000, 10000, 20000, 100000}
	workers := []int{2, 4, 10, 12, 25, 50}

	for i := 0; i < len(sizes); i++ {
		data := SequentialGenerateDataset(sizes[i])
		SeqAgg := SequentialAggregate(data)

		for j := 0; j < len(workers); j++ {
			ParAgg := ParallelAggregation_MapReduce(&data, workers[j])

			if SeqAgg == nil || ParAgg == nil {
				return false, "nilResult" + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			} else if SeqAgg != ParAgg {
				return false, "DiffResultAggregate: " + strconv.Itoa(sizes[i]) + ":" + strconv.Itoa(workers[j])
			}

		}
	}
	return true, ""

}
