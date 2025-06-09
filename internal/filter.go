package additionalFunctions

import (
	"sync"
)

func More(a, b any) bool {
	switch a.(type) {
	case int64:
		a_val, _ := a.(int64)
		b_val, _ := b.(int64)
		return a_val > b_val

	case float32:
		a_val, _ := a.(float32)
		b_val, _ := b.(float32)
		return a_val > b_val

	default:
		return false
	}
}

func Less(a, b any) bool {
	switch a.(type) {
	case int64:
		a_val, _ := a.(int64)
		b_val, _ := b.(int64)
		return a_val < b_val

	case float32:
		a_val, _ := a.(float32)
		b_val, _ := b.(float32)
		return a_val < b_val

	default:
		return false
	}
}

func Equal(a, b any) bool {
	switch a.(type) {
	case int64:
		a_val, _ := a.(int64)
		b_val, _ := b.(int64)
		return a_val > b_val

	case float32:
		a_val, _ := a.(float32)
		b_val, _ := b.(float32)
		return a_val > b_val

	default:
		a_val, _ := a.(bool)
		b_val, _ := b.(bool)
		return a_val == b_val
	}
}

func MoreEqual(a, b any) bool {
	switch a.(type) {
	case int64:
		a_val, _ := a.(int64)
		b_val, _ := b.(int64)
		return a_val >= b_val

	case float32:
		a_val, _ := a.(float32)
		b_val, _ := b.(float32)
		return a_val >= b_val

	default:
		return false
	}
}

func LessEqual(a, b any) bool {

	switch a.(type) {
	case int64:
		a_val, _ := a.(int64)
		b_val, _ := b.(int64)
		return a_val <= b_val

	case float32:
		a_val, _ := a.(float32)
		b_val, _ := b.(float32)
		return a_val <= b_val

	default:
		return false
	}
}

func ParallelFilter(records []int, workers, compare_value int) []int {
	if len(records) == 0 || workers <= 0 {
		return nil
	}
	if workers > len(records) {
		workers = len(records)
	}

	chunkSize := (len(records) + workers - 1) / workers
	chunks := make([][]int, workers)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= len(records) {
			start = len(records)
			end = start
		}
		if end > len(records) {
			end = len(records)
		}
		wg.Add(1)
		go func(i, start, end int) {
			defer wg.Done()
			size := end - start
			if size <= 0 {
				chunks[i] = nil
				return
			}
			part := make([]int, 0, size)
			for _, val := range records[start:end] {
				if val > compare_value {
					part = append(part, val)
				}
			}
			chunks[i] = part
		}(i, start, end)
	}

	wg.Wait()

	total := 0
	for _, c := range chunks {
		total += len(c)
	}
	res := make([]int, 0, total)
	for _, c := range chunks {
		res = append(res, c...)
	}
	return res
}

func SequentialFilter(records []int, compare_value int) []int {

	result := make([]int, 0)
	for _, val := range records {
		if val > compare_value {
			result = append(result, val)
		}
	}
	return result
}
