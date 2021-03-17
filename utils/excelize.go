package utils

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type XLS struct {
	FileName string
	Rows     [][]string
	Cols     [][]string
}

func (x *XLS) ParsingXLS() error {
	f, err := excelize.OpenFile(x.FileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Get all the rows in the Sheet1.
	x.Rows, err = f.GetRows("Localizable.strings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// for _, row := range x.Rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	// Get all the cols in the Sheet1.
	x.Cols, err = f.GetCols("Localizable.strings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// for _, col := range x.Cols {
	// 	for _, rowCell := range col {
	// 		fmt.Print(rowCell, "\n")
	// 	}
	// 	fmt.Println()
	// }

	return err
}

func (x *XLS) SaveXLS(axis string, newCol []string) error {
	f, err := excelize.OpenFile(x.FileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a new sheet.
	// index := f.NewSheet("Localizable.strings")
	index := f.GetSheetIndex("Localizable.strings")
	for i, rowCell := range newCol {
		f.SetCellStr("Localizable.strings", axis+strconv.Itoa(i+1), rowCell)
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
	return err
}

func NewXLS(name string) XLS {
	xls := XLS{
		FileName: name,
		Rows:     nil,
		Cols:     nil,
	}
	return xls
}
