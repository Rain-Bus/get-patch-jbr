package util

import (
	"fmt"
	"github.com/fatih/color"
	"strconv"
	"strings"
)

type SerialTable struct {
	TableHeader []string
	EleWidth    []int
	Elements    map[string][]string
	containZero bool
	tableHeight int
	tableWidth  int
}

// GetYNConfirm get yes/no confirm message from command line
// @param	confirmMsg		this is the message that ask user
// @param	defaultVal		when user press enter with blank, this is the default value, if no default，give it ""
func GetYNConfirm(confirmMsg string, defaultVal string /*"y" or "n" or ""*/) bool /*y: true, n: false*/ {

	var defaultTipFlag string

	yesMap := map[string]string{"y": "y", "yes": "y"}
	noMap := map[string]string{"n": "n", "no": "n"}

	defaultVal = strings.ToLower(defaultVal)
	switch defaultVal {
	case "y":
		defaultTipFlag = " [Y/n]"
	case "n":
		defaultTipFlag = " [y/N]"
	default:
		defaultTipFlag = " [y/n]"
	}

	color.Green("%s %s:", confirmMsg, defaultTipFlag)
	for true {
		var confirmStr string

		_, _ = fmt.Scanln(&confirmStr)
		confirmStr = strings.ToLower(strings.TrimSpace(confirmStr))

		if yesMap[confirmStr] != "" {
			return true
		}
		if noMap[confirmStr] != "" {
			return false
		}
		if confirmStr == "" && defaultVal == "y" {
			return true
		}
		if confirmStr == "" && defaultVal == "n" {
			return false
		}
		color.Yellow("Please retry input %s: ", defaultTipFlag)
	}

	return false

}

func (serialTable *SerialTable) ShowSerialTable() {
	serialTable.validateSerialTable()
	tableWidth := len(serialTable.TableHeader)
	tableHeight := 0
	for _, elements := range serialTable.Elements {
		tableHeight = len(elements)
	}
	serialColWidth := GetIntegerLength(tableHeight)

	formatter := "%" + strconv.Itoa(serialColWidth) + "s  "
	for colIndex := 0; colIndex < tableWidth; colIndex++ {
		formatter += "%-" + strconv.Itoa(serialTable.EleWidth[colIndex]) + "s"
	}
	formatter += "\n"

	serialTable.printTableHeader(formatter)
	serialTable.printTableBody(formatter, tableHeight)

}

func (serialTable *SerialTable) printTableHeader(formatter string) {
	tableHeader := []string{""}
	tableHeader = append(tableHeader, serialTable.TableHeader...)
	_, _ = color.New(color.FgBlue, color.Bold).Printf(formatter, ConvertStrArr2Inter(tableHeader)...)
}

func (serialTable *SerialTable) printTableBody(formatter string, tableHeight int) {
	var tableRow []string

	for row := 0; row < tableHeight; row++ {
		if serialTable.containZero {
			tableRow = append(tableRow, strconv.Itoa(row))
		} else {
			tableRow = append(tableRow, strconv.Itoa(row+1))
		}
		for _, header := range serialTable.TableHeader {
			tableRow = append(tableRow, serialTable.Elements[header][row])
		}
		color.Cyan(formatter, ConvertStrArr2Inter(tableRow)...)
		tableRow = tableRow[0:0]
	}
}

func (serialTable *SerialTable) validateSerialTable() {
	tableWidth := len(serialTable.TableHeader)
	// Judge map contain all header
	var eleHeaders []string
	for eleHeader, _ := range serialTable.Elements {
		eleHeaders = append(eleHeaders, eleHeader)
	}

	var tableHeader []string
	copy(tableHeader, serialTable.TableHeader)
	if len(DiffStrArr(tableHeader, eleHeaders)) != 0 {
		panic("The elements not match to table header")
	}

	if serialTable.EleWidth != nil && len(serialTable.EleWidth) != tableWidth {
		panic("The element width number is not fit to the number of table header")
	}

	lastLen := -1
	for header, elements := range serialTable.Elements {
		if lastLen == -1 {
			lastLen = len(elements)
		}
		if lastLen != len(elements) {
			panic("The column '" + header + "' length not match to other columns")
		}
	}

	serialTable.tableHeight = lastLen
	serialTable.tableWidth = tableWidth

}

// GetNumConfirm
// return the index of the selected element not the selected number
func (serialTable *SerialTable) GetNumConfirm(confirmMsg string) int {
	var minRange, maxRange int
	if serialTable.containZero {
		minRange = 0
	} else {
		minRange = 1
	}
	maxRange = minRange + serialTable.tableHeight - 1
	rangeStr := "(" + strconv.Itoa(minRange) + "-" + strconv.Itoa(maxRange) + ")"

	color.Green("%s %s:", confirmMsg, rangeStr)

	var selectNumStr string
	for true {
		_, err := fmt.Scanln(&selectNumStr)
		selectNum, err := strconv.Atoi(selectNumStr)
		if err != nil || selectNum < minRange || selectNum > maxRange {
			color.Yellow("Please retry input %s:", rangeStr)
			continue
		}
		if serialTable.containZero {
			return selectNum
		}
		return selectNum - 1
	}

	return minRange
}

func NewSerialTable(containZero bool) *SerialTable {
	return &SerialTable{containZero: containZero}
}
