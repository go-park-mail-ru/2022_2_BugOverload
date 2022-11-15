package sqltools

import (
	"database/sql"
	"time"
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

func NewSQLNNullDate(date string, format string) sql.NullTime {
	dateTime, err := time.Parse(format, date)
	if err != nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  dateTime,
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
