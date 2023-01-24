package main

import (
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"

	"github.com/xitongsys/parquet-go-source/local"
	"log"
)

type Student struct {
	Name string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Age  int32  `parquet:"name=age, type=INT32"`
	Id   int64  `parquet:"name=id, type=INT64"`
}

func main() {
	var err error
	fw, err := local.NewLocalFileWriter("flat.parquet")
	if err != nil {
		log.Println("Can't create local file", err)
		return
	}

	//write
	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
	if err != nil {
		log.Println("Can't create parquet writer", err)
		return
	}

	pw.RowGroupSize = 128 * 1024 * 1024 //128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	num := 100
	for i := 0; i < num; i++ {
		stu := Student{
			Name: "StudentName",
			Age:  int32(20 + i%5),
			Id:   int64(i),
		}
		if err = pw.Write(stu); err != nil {
			log.Println("Write error", err)
		}
	}
	if err = pw.WriteStop(); err != nil {
		log.Println("WriteStop error", err)
		return
	}
	log.Println("Write Finished")
	fw.Close()

	///read
	fr, err := local.NewLocalFileReader("flat.parquet")
	if err != nil {
		log.Println("Can't open file")
		return
	}

	pr, err := reader.NewParquetReader(fr, new(Student), 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return
	}
	num = int(pr.GetNumRows())
	for i := 0; i < num/10; i++ {
		if i%2 == 0 {
			pr.SkipRows(10) //skip 10 rows
			continue
		}
		stus := make([]Student, 10) //read 10 rows
		if err = pr.Read(&stus); err != nil {
			log.Println("Read error", err)
		}
		log.Println(stus)
	}

	pr.ReadStop()
	fr.Close()

}
