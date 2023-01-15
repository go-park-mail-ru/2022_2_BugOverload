package devpkg

const (
	TypeDataReal   = "real"
	TypeDataRandom = "random"

	TypeSerial = "serial"
	TypeFilm   = "film"

	// 65535
	MaxInsertValuesPostgreSQL = 65535

	MaxInsertValuesSQL = MaxInsertValuesPostgreSQL - 100

	TypeReviewPositive = "positive"
	TypeReviewNegative = "negative"
	TypeReviewNeutral  = "neutral"
)
