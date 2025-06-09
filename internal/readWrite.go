package additionalFunctions

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SequentialRead(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}

	var numbers []int
	for _, row := range records {
		for _, col := range row {
			num, err := strconv.Atoi(col)
			if err != nil {
				return nil
			}
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func SaveEfficiency(Efficiency [][]EfficiencyResult, file string, sizes, workersArr []int) error {

	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open/create file: %v", err)
	}
	defer f.Close()

	for i := 0; i < len(sizes); i++ {
		dataSize := sizes[i]

		index := 0
		for j := 0; j < len(workersArr); j++ {
			workers := workersArr[j]

			currentEfficiency := Efficiency[i][index]

			_, err = fmt.Fprintf(f, "size: %d, g: %d, threads: %d, timeS: %.5f, timeP: %.5f, SpeedUp: %.2f, Eff: %.2f, trueg: %d\n",
				dataSize, workers, currentEfficiency.threads,
				currentEfficiency.timeSeq, currentEfficiency.timePar,
				currentEfficiency.speedUp, currentEfficiency.efficiency,
				currentEfficiency.trueCountGoroutines,
			)
			if err != nil {
				return fmt.Errorf("failed to write data: %v", err)
			}

		}
	}

	return nil
}

func SequentialSaveCSV(numbers []int, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, num := range numbers {

		row := []string{strconv.Itoa(num)}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("ошибка записи в CSV: %v", err)
		}
	}

	return nil
}
