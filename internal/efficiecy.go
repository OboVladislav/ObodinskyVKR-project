package additionalFunctions

import (
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

type EfficiencyResult struct {
	size                int
	g                   int
	timeSeq             float64
	timePar             float64
	speedUp             float64
	efficiency          float64
	trueCountGoroutines int64
	threads             int
}

func GenerateTestDatasets(sizes []int) []string {
	files := []string{}
	for i := 0; i < len(sizes); i++ {
		dataSize := sizes[i]
		fmt.Println(dataSize)
		currentDataset := SequentialGenerateDataset(dataSize)
		fmt.Println(len(currentDataset))
		SequentialSaveCSV(currentDataset, "datasets/TestData"+strconv.Itoa(dataSize)+"_4.csv")
		files = append(files, "datasets/TestData"+strconv.Itoa(dataSize)+".csv")
	}
	return files
}

func PrintEfficiency(Efficiency [][]EfficiencyResult, sizes, workersArr []int) {
	for i := 0; i < len(sizes); i++ {
		dataSize := sizes[i]
		fmt.Println("Size:", dataSize)

		index := 0
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]

			for k := 1; k <= runtime.NumCPU(); k++ {
				currentEfficiency := Efficiency[i][index]

				fmt.Printf("size: %d, g: %d, threads: %d, timeS: %.5f, timeP: %.5f, SpeedUp: %.2f, Eff: %.2f, trueg: %d\n",
					dataSize, workers, currentEfficiency.threads,
					currentEfficiency.timeSeq, currentEfficiency.timePar,
					currentEfficiency.speedUp, currentEfficiency.efficiency,
					currentEfficiency.trueCountGoroutines,
				)

				index++
			}
		}
	}
}

func GetEfficiencyGenerate(sizes, workersArr []int) [][]EfficiencyResult {
	// var timeSeq time.Duration
	result := [][]EfficiencyResult{}
	for i := 0; i < len(sizes); i++ {
		dataSize := sizes[i]

		startSeq := time.Now()
		// SequentialGenerateDataset(dataSize)
		ParallelGenerateDataset(dataSize, 1)
		timeSeq := time.Since(startSeq)

		EfficiencyPerWorker := []EfficiencyResult{}
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]
			for k := 1; k <= runtime.NumCPU(); k++ {

				// if k == 12 {
				runtime.GOMAXPROCS(k)

				startPar := time.Now()
				ParallelGenerateDataset(dataSize, workers)
				timePar := time.Since(startPar)
				// if k == 1 {
				// 	timeSeq = timePar
				// }

				speedUp := float64(timeSeq) / float64(timePar)
				efficiency := speedUp / float64(k)

				EfficiencyPerWorker = append(EfficiencyPerWorker, EfficiencyResult{dataSize, workers, timeSeq.Seconds(), timePar.Seconds(), speedUp, efficiency, int64(workers), k})

				// }

			}
		}
		result = append(result, EfficiencyPerWorker)
	}
	return result
}

func GetEfficiencyMergeSort(files []string, workersArr []int) [][]EfficiencyResult {

	// var timeSeq time.Duration
	result := [][]EfficiencyResult{}
	for i := 0; i < len(files); i++ {
		file := files[i]
		currentDataNotSorted := SequentialRead(file)

		currentDataSeq := make([]int, len(currentDataNotSorted))
		copy(currentDataSeq, currentDataNotSorted)

		startSeq := time.Now()
		ParallelMergeSort(currentDataSeq, 1)
		timeSeq := time.Since(startSeq)

		EfficiencyPerWorker := []EfficiencyResult{}
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]

			for k := 1; k <= runtime.NumCPU(); k++ {

				// if k == 1 {
				runtime.GOMAXPROCS(k)

				currentDataPar := make([]int, len(currentDataNotSorted))
				copy(currentDataPar, currentDataNotSorted)

				// f, err := os.Create("cpu.pprof")
				// if err != nil {
				// 	panic(err)
				// }
				// defer f.Close()
				// if err := pprof.StartCPUProfile(f); err != nil {
				// 	panic(err)
				// }
				// defer pprof.StopCPUProfile()

				startPar := time.Now()
				ParallelMergeSort(currentDataPar, workers)
				timePar := time.Since(startPar)

				// if k == 1 {
				// 	timeSeq = timePar
				// }

				speedUp := float64(timeSeq) / float64(timePar)
				efficiency := speedUp / float64(k)

				EfficiencyPerWorker = append(EfficiencyPerWorker, EfficiencyResult{len(currentDataNotSorted), workers, timeSeq.Seconds(), timePar.Seconds(), speedUp, efficiency, pCounter, k})
				currentDataPar = nil

				// }

			}
		}

		result = append(result, EfficiencyPerWorker)

		// Очистка ссылок и сбор мусора
		currentDataNotSorted = nil
		currentDataSeq = nil
		runtime.GC()
	}
	return result
}

func GetEfficiencyQuickSort(files []string, workersArr []int) [][]EfficiencyResult {

	var timeSeq time.Duration
	result := [][]EfficiencyResult{}
	for i := 0; i < len(files); i++ {
		file := files[i]
		currentDataNotSorted := SequentialRead(file)

		currentDataSeq := make([]int, len(currentDataNotSorted))
		copy(currentDataSeq, currentDataNotSorted)

		// startSeq := time.Now()
		// SequentialQuicksort(currentDataSeq, 0, len(currentDataSeq)-1)
		// timeSeq := time.Since(startSeq)

		EfficiencyPerWorker := []EfficiencyResult{}
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]

			for k := 1; k <= runtime.NumCPU(); k++ {

				if k == 1 {
					runtime.GOMAXPROCS(k)

					currentDataPar := make([]int, len(currentDataNotSorted))
					copy(currentDataPar, currentDataNotSorted)
					atomic.StoreInt64(&pCounterQ, 0)

					startPar := time.Now()
					ParallelQuicksort(currentDataPar, workers)
					timePar := time.Since(startPar)

					if k == 1 {
						timeSeq = timePar
					}

					speedUp := float64(timeSeq) / float64(timePar)
					efficiency := speedUp / float64(k)

					EfficiencyPerWorker = append(EfficiencyPerWorker, EfficiencyResult{len(currentDataSeq), workers, timeSeq.Seconds(), timePar.Seconds(), speedUp, efficiency, pCounterQ, k})

				}

			}
		}

		result = append(result, EfficiencyPerWorker)
	}
	return result
}

func GetEfficiencyFilter(files []string, workersArr []int, compare_value int) [][]EfficiencyResult {
	result := [][]EfficiencyResult{}
	var timeSeq time.Duration
	for _, file := range files {
		func() {
			currentData := SequentialRead(file)

			metrics := ParallelAggregation_MapReduce(&currentData, 1000).(struct {
				Avg   float32
				Max   int
				Min   int
				Count int
			})
			compare_value = int(metrics.Avg)
			// startSeq := time.Now()
			// SequentialFilter(currentData, compare_value)
			// timeSeq := time.Since(startSeq)

			efficiencyPerWorker := make([]EfficiencyResult, 0, len(workersArr)*runtime.NumCPU())
			for _, workers := range workersArr {
				for k := 1; k <= runtime.NumCPU(); k++ {

					if k == 4 || k == 1 {
						runtime.GOMAXPROCS(k)

						startPar := time.Now()
						ParallelFilter(currentData, workers, compare_value)
						timePar := time.Since(startPar)

						if k == 1 {
							timeSeq = timePar
						}

						speedUp := float64(timeSeq) / float64(timePar)
						efficiency := speedUp / float64(k)

						efficiencyPerWorker = append(efficiencyPerWorker, EfficiencyResult{
							size:                len(currentData),
							g:                   workers,
							timeSeq:             timeSeq.Seconds(),
							timePar:             timePar.Seconds(),
							speedUp:             speedUp,
							efficiency:          efficiency,
							trueCountGoroutines: int64(workers),
							threads:             runtime.GOMAXPROCS(0),
						})
					}

				}
			}
			result = append(result, efficiencyPerWorker)
		}()

	}
	return result
}
func GetEfficiencyAggregate(files []string, workersArr []int) [][]EfficiencyResult {

	// var timeSeq time.Duration
	result := [][]EfficiencyResult{}
	for i := 0; i < len(files); i++ {
		file := files[i]
		currentData := SequentialRead(file)
		startSeq := time.Now()

		ParallelAggregation_MapReduce(&currentData, 1)
		timeSeq := time.Since(startSeq)

		// startSeq := time.Now()
		// SequentialAggregate(currentData)
		// timeSeq := time.Since(startSeq)

		EfficiencyPerWorker := []EfficiencyResult{}
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]
			for k := 1; k <= runtime.NumCPU(); k++ {

				// if k == 5 || k == 1 {
				runtime.GOMAXPROCS(k)

				startPar := time.Now()
				ParallelAggregation_MapReduce(&currentData, workers)
				timePar := time.Since(startPar)

				// if k == 1 {
				// 	timeSeq = timePar
				// }

				speedUp := float64(timeSeq) / float64(timePar)
				efficiency := speedUp / float64(k)

				EfficiencyPerWorker = append(EfficiencyPerWorker, EfficiencyResult{len(currentData), workers, timeSeq.Seconds(), timePar.Seconds(), speedUp, efficiency, int64(workers), runtime.GOMAXPROCS(0)})

				// }

			}
		}
		result = append(result, EfficiencyPerWorker)
	}
	return result
}
