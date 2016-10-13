package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aswjh/excel"
)

const (
	// 报表名称
	REPORT_NAME = "report.xls"
	// timeout to delete
	DELETE_TIME = 60
)

var (
	reportXlsPath string
)

func main() {
	currentDir, xlsFiles, err := scanXlsFiles()
	if err != nil {
		log.Printf("scan xls files error: %v\n", err)
		return
	}

	fmt.Println("********These files will be processed")
	for _, f := range xlsFiles {
		fmt.Println(f)
	}
	fmt.Printf("* Total %d\n", len(xlsFiles))

	if len(xlsFiles) <= 0 {
		fmt.Println("Please move this binary aside xls files")
		return
	}

	fmt.Print("Is is OK(y/N)?")
	answer, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if strings.TrimSpace(answer) != "Y" && strings.TrimSpace(answer) != "y" {
		fmt.Println("Cancelled!")
		return
	}

	t1 := time.Now()

	fmt.Println("- Init report table...")
	reportXlsPath = filepath.Join(currentDir, REPORT_NAME)
	initReportXls(reportXlsPath)

	for _, f := range xlsFiles {
		err := process(filepath.Join(currentDir, f))
		if err != nil {
			log.Printf("process file error: %v\n", err)
			return
		}
	}
	fmt.Printf("* Time spent total: %vs\n", time.Since(t1).Seconds())
}

func scanXlsFiles() (currentDir string, xlsFiles []string, err error) {
	currentDir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("get current dir error: %v\n", err)
		return
	}

	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		log.Printf("list files in dir %s error: %v\n", currentDir, err)
		return
	}

	fmt.Println("- Scanning current dir...")
	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("skip dir \"%s\"\n", file.Name())
			continue
		}
		if !strings.Contains(file.Name(), "xls") {
			fmt.Printf("skip file \"%s\"\n", file.Name())
			continue
		}
		if strings.Contains(file.Name(), "xlsx") {
			fmt.Printf("skip file \"%s\", xlsx not supported\n", file.Name())
			continue
		}
		if file.Name() == REPORT_NAME {
			fmt.Printf("skip report file \"%s\"\n", file.Name())
			continue
		}

		fmt.Printf("add file \"%s\"\n", file.Name())
		xlsFiles = append(xlsFiles, file.Name())
	}
	return
}

func initReportXls(path string) (err error) {
	// check if report exist
	if _, err = os.Stat(path); err == nil {
		for i := 0; i < DELETE_TIME; i++ {
			fmt.Printf("\rfile \"%s\" already exist, will DELETE it in %ds, backup now", path, DELETE_TIME-i)
			time.Sleep(1 * time.Second)
		}

		fmt.Println("")

		err = os.Remove(path)
		if err != nil {
			fmt.Printf("delete file %s error: %v\n", path, err)
		}
		fmt.Printf("file %s deleted\n", path)
	}

	// write to excel
	option := excel.Option{"Visible": false, "DisplayAlerts": false}
	resultXls, err := excel.New(option)
	if err != nil {
		log.Printf("new excel error: %v\n", err)
		return
	}
	defer resultXls.Quit()

	// init first line
	sheet, _ := resultXls.Sheet(1)
	sheet.PutCell(1, 1, "号段")
	sheet.PutCell(1, 2, "序号")
	sheet.PutCell(1, 3, fmt.Sprintf("频率最高%d个", N_CODE_TOP))
	sheet.PutCell(1, 4, "总次数")
	for i := HOUR_BEGIN; i < HOUR_END; i++ {
		sheet.PutCell(1, i-HOUR_BEGIN+5, fmt.Sprintf("%d~%d\t", i, i+1))
	}
	errArr := resultXls.SaveAs(path)
	if len(errArr) > 0 {
		if len(errArr) == 1 && errArr[0] == nil {
			return
		}
		log.Printf("save result xls error: %v\n", errArr)
		return
	}
	return
}
