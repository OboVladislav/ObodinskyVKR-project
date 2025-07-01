package main

import (
	. "vkr-req/internal"
)

func main() {

	TestName := "threads"
	// TestName := "data_values"
	if TestName == "threads" {

		sizes := []int{100000000}
		workers := []int{12, 500, 1000}
		tests := []string{"datasets/TestData100000000.csv",
			"datasets/TestData100000000_1.csv", "datasets/TestData100000000_2.csv",
			"datasets/TestData100000000_3.csv", "datasets/TestData100000000_4.csv"}

		// generate
		eff_gen := GetEfficiencyGenerate(sizes, workers)
		SaveEfficiency(eff_gen, "tripletest/GenTest.txt", sizes, workers)

		// filter
		eff_filter := GetEfficiencyFilter(tests, workers, 1)
		SaveEfficiency(eff_filter, "tripletest/FilterTest.txt", sizes, workers)

		// MergeSort
		eff_sortM := GetEfficiencyMergeSort(tests, workers)
		SaveEfficiency(eff_sortM, "tripletest/MSTest.txt", sizes, workers)

		// QuickSort
		eff_sort := GetEfficiencyQuickSort(tests, workers)
		SaveEfficiency(eff_sort, "tripletest/QSTest.txt", sizes, workers)

		// aggregate
		eff_agg := GetEfficiencyAggregate(tests, workers)
		SaveEfficiency(eff_agg, "tripletest/AggTest.txt", sizes, workers)

	} else if TestName == "data_values" {

		sizes := []int{4000000, 40000000, 400000000}
		tests := []string{"datasets/TestData4000000.csv",
			"datasets/TestData40000000.csv",
			"datasets/TestData400000000.csv",
		}
		// generate
		workers := []int{12}
		eff_gen := GetEfficiencyGenerate(sizes, workers)
		SaveEfficiency(eff_gen, "tripletest/GenTest.txt", sizes, workers)

		// filter
		workers = []int{500}
		eff_filter := GetEfficiencyFilter(tests, workers, 1)
		SaveEfficiency(eff_filter, "tripletest/FilterTest.txt", sizes, workers)

		// MergeSort
		workers = []int{1000}
		eff_sortM := GetEfficiencyMergeSort(tests, workers)
		SaveEfficiency(eff_sortM, "tripletest/MSTest.txt", sizes, workers)

		// QuickSort
		workers = []int{1000}
		eff_sort := GetEfficiencyQuickSort(tests, workers)
		SaveEfficiency(eff_sort, "tripletest/QSTest.txt", sizes, workers)

		// aggregate
		workers = []int{1000}
		eff_agg := GetEfficiencyAggregate(tests, workers)
		SaveEfficiency(eff_agg, "tripletest/AggTest.txt", sizes, workers)
	}

}
