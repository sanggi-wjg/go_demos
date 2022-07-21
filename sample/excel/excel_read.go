package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ExcelReadMain() {

	f, err := excelize.OpenFile("sample/excel/sample_read.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	sheet := sheets[1]

	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
	}

	for i, row := range rows {
		fmt.Println(i, row)
	}
}

/*
0  [ 날짜 번호  지역 무게 가격 비고]
1  [ 08월 01일 NO1  서울 0.35 4.5 1]
2  [ 08월 02일 NO2  경기 0.3 4.5 2]
3  [ 08월 03일 NO3  부산 0.22 7.5 3]
4  [ 08월 04일 NO4  경남 0.35 4.5 4]
5  [ 08월 05일 NO5  경북 0.2 11.5 5]
6  [ 08월 06일 NO6  인천 0.34 4.5 6]
7  [ 08월 07일 NO7  이천 0.78 11.5 7]
8  [ 08월 08일 NO8  장호원 1.2 5.5 8]
9  [ 08월 09일 NO9  이황리 0.23 4.5 9]
10 [ 08월 10일 NO10  중원 1.8 5.5 10]
*/
