package main

import (
	. "vkr-req/internal"
)

func main() {

	// CheckGenerate()
	// CheckFilter()
	// CheckMergeSort()
	// CheckQuickSort()
	// CheckAggregate()

	// // AnalysisSel()
	// sizes := []int{50000, 100000, 500000, 1000000, 5000000, 10000000, 50000000, 100000000}
	sizes := []int{5000000, 10000000, 15000000, 20000000, 25000000, 75000000, 100000000}
	// // sizes := []int{10000000, 100000000}
	workers := []int{500}
	// workers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 20, 100, 250, 500, 1000}
	// // workers := []int{12, 500, 1000}
	// GenerateTestDatasets(sizes)

	// tests := []string{"datasets/TestData50000.csv",
	// 	"datasets/TestData100000.csv",
	// 	"datasets/TestData500000.csv", "datasets/TestData1000000.csv", "datasets/TestData5000000.csv",
	// 	"datasets/TestData10000000.csv", "datasets/TestData50000000.csv", "datasets/TestData100000000.csv"}

	// // // s := []int{50000, 100000, 500000}

	tests := []string{"datasets/TestData5000000.csv", "datasets/TestData10000000.csv", "datasets/TestData15000000.csv", "datasets/TestData20000000.csv",
		"datasets/TestData25000000.csv", "datasets/TestData75000000.csv", "datasets/TestData100000000.csv"}

	// // 	"datasets/TestData100000.csv",
	// // 	"datasets/TestData500000.csv"}

	// eff_gen := GetEfficiencyGenerate(sizes, workers)
	// SaveEfficiency(eff_gen, "results/Gen.txt", sizes, workers)

	// eff_sort := GetEfficiencyQuickSort(tests, workers)
	// // PrintEfficiency(eff_sort, sizes, workers)
	// SaveEfficiency(eff_sort, "results/QS.txt", sizes, workers)

	// eff_sortM := GetEfficiencyMergeSort(tests, workers)
	// SaveEfficiency(eff_sortM, "results/MS.txt", sizes, workers)
	// PrintEfficiency(eff_sortM, s, workers)

	// eff_filter := GetEfficiencyFilter(tests, workers, 345737)
	// SaveEfficiency(eff_filter, "results/Fil.txt", sizes, workers)

	eff_agg := GetEfficiencyAggregate(tests, workers)
	SaveEfficiency(eff_agg, "results/Agg.txt", sizes, workers)

}
