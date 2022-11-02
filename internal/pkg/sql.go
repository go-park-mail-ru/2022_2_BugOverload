package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func NewSQLNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func NewSQLNullInt32(i int) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}

func NewSQLNullFloat64(f float32) sql.NullFloat64 {
	if f == 0 {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: float64(f),
		Valid:   true,
	}
}

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

func CreateStatement(query string, countInserts int) (string, int) {
	countAttributes := strings.Count(query, ",") + 1

	placeholders := CreatePlaceholders(countAttributes, countInserts)

	insertStatement := fmt.Sprintf("%s %s", query, placeholders)

	return insertStatement, countAttributes
}

func InsertBatch(ctx context.Context, db *sql.DB, query string, values []interface{}) (sql.Result, error) {
	rows, err := db.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, fmt.Errorf("InsertBatch: [%w] when inserting row into [%s] table", err, query)
	}

	return rows, nil
}
