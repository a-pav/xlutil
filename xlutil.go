package xlutil

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

func ExcelColumnNo(name string) int {
	var (
		number = 0
		pow    = 1
	)
	for i := len(name) - 1; i >= 0; i-- {
		number += int(name[i]-'A'+1) * pow
		pow *= 26
	}

	return number - 1
}

func ExcelColumnName(number int) string {
	number = number + 1
	var (
		column string
		mod    int
	)
	for number > 0 {
		mod = (number - 1) % 26
		column = string(65+mod) + column
		number = (number - mod) / 26
	}

	return column
}

func Axis(row, col int) (axis string) {
	return ExcelColumnName(col) + strconv.Itoa(row+1)
}

func Pretty(debugNote, x interface{}, indentStyle ...string) {
	var indent string
	if len(indentStyle) == 0 {
		indent = "\t|" // default indent style
	} else {
		indent = strings.Join(indentStyle, "")
	}

	j, err := json.MarshalIndent(x, "", indent)
	if err != nil {
		log.Println(err)
	}
	log.Println(debugNote, string(j))
}
