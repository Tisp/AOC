package utils

import (
	"fmt"
	"os"
)

func ReadSampleByDay(sample, day int) string {
	file_path := fmt.Sprintf("samples/sample-%d-day%d.txt", sample, day)
	data, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
