package csv_util

import (
	"bufio"
	"encoding/csv"
	"os"
)

func Write(path string) {
	// 파일 생성
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Flush()
}
