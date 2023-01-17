package sqltools

import (
	"fmt"
	"strings"
)

func CreatePlaceholders(countAttributes int, countValues int) string {
	values := make([]string, countAttributes*countValues)

	for i := 0; i < countAttributes*countValues; i++ {
		values[i] = fmt.Sprintf("$%d", i+1)
	}

	valuesRow := make([]string, countValues)

	for i := 0; i < countValues; i++ {
		valuesRow[i] = "(" + strings.Join(values[i*countAttributes:countAttributes*(i+1)], ",") + ")"
	}

	return strings.Join(valuesRow, ",\n")
}

func CreateFullQuery(query string, countInserts int, countAttributes int) string {
	placeholders := CreatePlaceholders(countAttributes, countInserts)

	insertStatement := fmt.Sprintf("%s %s", query, placeholders)

	return insertStatement
}
