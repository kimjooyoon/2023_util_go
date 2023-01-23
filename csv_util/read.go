package csv_util

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Read(path string) {
	file, err1 := os.Open(path)
	if err1 != nil {
		log.Panicf("%v", err1)
	}

	rdr := csv.NewReader(bufio.NewReader(file))

	// csv 내용 모두 읽기
	rows, err2 := rdr.ReadAll()
	if err2 != nil {
		log.Panicf("%v", err2)
	}

	// 행,열 읽기
	for i, row := range rows {
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}
}
