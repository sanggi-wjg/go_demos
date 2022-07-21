package excel

import (
	"github.com/xuri/excelize/v2"
)

func ExcelWriteMain() {
	sheetName := "NewSheet"
	f := excelize.NewFile()

	index := f.NewSheet(sheetName)

	f.SetCellValue(sheetName, "B2", "Hello")
	f.SetCellValue(sheetName, "B3", "World")

	f.SetActiveSheet(index)

	if err := f.SaveAs("NewExcel.xlsx"); err != nil {
		panic(err)
	}
}
