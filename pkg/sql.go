package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
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

func SendQuery(db *sql.DB, timeout int, insertStatement string, target string, values []interface{}) (*sql.Stmt, *sql.Rows, context.CancelFunc, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	stmt, err := db.PrepareContext(ctx, insertStatement)
	if err != nil {
		logrus.Errorf("Error [%s] when preparing SQL statement in [%s]", err, target)
		return nil, nil, cancelFunc, err
	}

	rows, err := stmt.QueryContext(ctx, values...)
	if err != nil {
		logrus.Errorf("Error [%s] when inserting row into [%s] table", err, target)
		return stmt, nil, cancelFunc, err
	}

	return stmt, rows, cancelFunc, nil
}
