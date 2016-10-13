package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/aswjh/excel"
)

const (
	// 时间 列号
	COL_TIME = 1
	// 业务代码 列号
	COL_CODE = 3

	// 数据首行 行号
	ROW_DATA = 5

	// 最大业务代码个数
	N_CODE_TOP = 30

	// 统计几点开始
	HOUR_BEGIN = 9
	// 统计几点结束
	HOUR_END = 15
)

// for offset
var writeCount int

func process(file string) (err error) {
	t1 := time.Now()

	fmt.Printf("======>> Processing %s...\n", file)
	option := excel.Option{"Visible": false, "DisplayAlerts": true}
	mso, err := excel.Open(file, option)
	if err != nil {
		log.Printf("open excel error: %v\n", err)
		return
	}
	defer mso.Quit()
	fmt.Printf("* Excel version: %v\n", mso.Version)

	fmt.Printf("* Workbooks: %d\n", mso.CountWorkBooks())
	fmt.Printf("* Sheets: %d\n", mso.CountSheets())

	for i, sheet := range mso.Sheets() {
		if !strings.Contains(sheet.Name(), "号段") {
			fmt.Printf("- Skip sheet %d: %s due to sheet name\n", i, sheet.Name())
			continue
		}
		title, _ := sheet.GetCell(1, 1)
		if !strings.Contains(excel.String(title), "办理明细") {
			fmt.Printf("- Skip sheet %d: %s due to title\n", i, excel.String(title))
			continue
		}

		fmt.Printf("- Process sheet %d: %s\n", i, sheet.Name())

		// get row count
		fmt.Println("- Counting rows...")
		row := 0
		for c := 1; c < math.MaxInt32; c++ {
			cell, _ := sheet.GetCell(c, COL_TIME)
			if excel.String(cell) == "已导出明细" {
				break
			}
			row++
		}
		fmt.Printf("* Rows: %d\n", row)

		fmt.Println("- Collecting codes...")
		// Ye Wu Dai Ma
		codeArr := []int{}

		// j: row
		for j := ROW_DATA; j <= row; j++ {
			// parse col 3 and append to codeCell
			codeCell, err := sheet.GetCell(j, COL_CODE)
			if err != nil {
				log.Printf("get code cell error: %v\n", err)
				continue
			}
			code, err := strconv.ParseInt(excel.String(codeCell), 10, 0)
			if err != nil {
				log.Printf("parse string to int error: %v\n", err)
				continue
			}
			codeArr = append(codeArr, int(code))
		}

		fmt.Println("- Sorting codes...")
		// get top 30 biggest number
		top := top(codeArr, N_CODE_TOP)
		fmt.Printf("* Top %d codes: %v\n", N_CODE_TOP, top)

		fmt.Println("- Generating frequency table...")
		//// Loop and count frequency
		// frequency table
		// row: Top 30 biggest code
		// col: 0-23h hour period
		var freqTable [N_CODE_TOP][24]int

		// j: row
		// k: col
		for j := ROW_DATA; j <= row; j++ {
			codeCell, err := sheet.GetCell(j, COL_CODE)
			if err != nil {
				log.Printf("get cell(%d,%d) error:%v\n", j, COL_CODE, err)
				continue
			}
			code, err := strconv.ParseInt(excel.String(codeCell), 10, 0)
			if err != nil {
				log.Printf("parse cell(%d,%d) to int error:%v\n", j, COL_CODE, err)
				continue
			}

			for k := 0; k < N_CODE_TOP; k++ {
				if int(code) == top[k] {
					timeCell, err := sheet.GetCell(j, COL_TIME)
					if err != nil {
						log.Printf("get cell(%d,%d) error:%v\n", j, COL_CODE, err)
						continue
					}
					// f is time like 0.375
					f, err := strconv.ParseFloat(excel.String(timeCell), 32)
					if err != nil {
						log.Printf("parse cell(%d,%d) to float error:%v\n", k, COL_TIME, err)
						continue
					}
					h := hour(float32(f))
					freqTable[k][h]++
					// fmt.Println(freqTable)
				}
			}
		}

		fmt.Println("- Calculating ratio table...")
		var trimedTable [N_CODE_TOP][HOUR_END - HOUR_BEGIN]int
		for j := 0; j < N_CODE_TOP; j++ {
			for k := 0; k < HOUR_END-HOUR_BEGIN; k++ {
				trimedTable[j][k] = freqTable[j][k+HOUR_BEGIN]
			}
		}

		var ratioTable [N_CODE_TOP][HOUR_END - HOUR_BEGIN]float32
		for j := 0; j < N_CODE_TOP; j++ {
			sum := 0
			for k := 0; k < HOUR_END-HOUR_BEGIN; k++ {
				sum += trimedTable[j][k]
			}
			for k := 0; k < HOUR_END-HOUR_BEGIN; k++ {
				if sum != 0 {
					ratioTable[j][k] = float32(trimedTable[j][k]) / float32(sum)
				}
			}
		}

		fmt.Printf("* Ratio table: %v\n", ratioTable)

		fmt.Printf("- Writing data to file %s...\n", reportXlsPath)

		numPeriod := strings.Trim(sheet.Name(), "号段")
		err := writeExcel(writeCount, trimedTable, ratioTable, top, reportXlsPath, numPeriod)
		if err != nil {
			log.Printf("write result to excel error: %v\n", err)
			continue
		}
		writeCount++
	}

	fmt.Printf("* Time spent: %vs\n", time.Since(t1).Seconds())
	return
}

// write to report
func writeExcel(n int, trimedTable [N_CODE_TOP][HOUR_END - HOUR_BEGIN]int, ratioTable [N_CODE_TOP][HOUR_END - HOUR_BEGIN]float32, top []int, fileName string, numPeriod string) (err error) {
	// write to excel
	option := excel.Option{"Visible": false, "DisplayAlerts": false}
	resultXls, err := excel.Open(fileName, option)
	if err != nil {
		log.Printf("new excel error: %v\n", err)
		return
	}
	defer resultXls.Quit()

	sheet, _ := resultXls.Sheet(1)

	offset := n * N_CODE_TOP
	for i := 0; i < N_CODE_TOP; i++ {
		sheet.PutCell(i+2+offset, 1, numPeriod)
		sheet.PutCell(i+2+offset, 2, i+1)
		sheet.PutCell(i+2+offset, 3, top[i])
		sum := 0
		for j := 0; j < HOUR_END-HOUR_BEGIN; j++ {
			sum += trimedTable[i][j]
		}
		sheet.PutCell(i+2+offset, 4, sum)
		for j := 0; j < HOUR_END-HOUR_BEGIN; j++ {
			sheet.PutCell(i+2+offset, j+5, fmt.Sprintf("%.2f%%", 100*ratioTable[i][j]))
		}
	}
	errArr := resultXls.SaveAs(fileName)
	if len(errArr) > 0 {
		if len(errArr) == 1 && errArr[0] == nil {
			return
		}
		log.Printf("save result xls error: %v\n", errArr)
		return
	}
	time.Sleep(3 * time.Second)
	return
}
