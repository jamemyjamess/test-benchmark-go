package main

import (
	"errors"
	"log"
	"math"
	"math/rand"
)

func main() {
	type DataWriteExcel struct {
		offset int
		limit  int
		data   []interface{}
	}
	done := make(chan bool)
	errCh := make(chan error)
	dataWriteExcelCh := make(chan DataWriteExcel)
	amountTyreRequest := 1390
	limitPerList := 500
	currentRowExcel := 2
	go func() {
		// จำนวนสัดส่วนทั้งหมด
		sec := int(math.Ceil(float64(amountTyreRequest) / float64(limitPerList)))
		for i := 0; i < sec; i++ {
			offset := i * limitPerList
			limit := limitPerList
			row := make([]interface{}, 50)
			for colID := 0; colID < 50; colID++ {
				row[colID] = rand.Intn(640000)
				if colID == 20 {
					errCh <- errors.New("Something went wrong.")
					close(dataWriteExcelCh)
					done <- true
					return
				}
			}
			dataWriteExcelCh <- DataWriteExcel{offset: offset, limit: limit, data: row}
		}
		done <- true
	}()

	go func() {
		for {
			select {
			case dataWriteExcel, ok := <-dataWriteExcelCh:
				if !ok {
					return
				}
				rowNo := currentRowExcel + dataWriteExcel.offset
				// cell, _ := excelize.CoordinatesToCellName(1, rowNo)
				// if err := streamWriter.SetRow(cell, dataWriteExcel.data, excelize.RowOpts{StyleID: textCenterTableStyle}); err != nil {
				// 	log.Println("stream writer error:", err.Error())
				// }
				log.Println("rowNo:", rowNo, "offset:", dataWriteExcel.offset, "limit:", dataWriteExcel.limit)
			}
		}
		// for dataWriteExcel := range dataWriteExcelCh {
		// 	rowNo := currentRowExcel + dataWriteExcel.offset
		// 	cell, _ := excelize.CoordinatesToCellName(1, rowNo)
		// 	if err := streamWriter.SetRow(cell, dataWriteExcel.data, excelize.RowOpts{StyleID: textCenterTableStyle}); err != nil {
		// 		log.Println("stream writer error:", err.Error())
		// 	}
		// }
	}()

	<-done
	if err := <-errCh; err != nil {
		log.Println(err.Error())
	}
}
