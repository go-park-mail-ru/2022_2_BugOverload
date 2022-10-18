package errors

type errClassifier interface {
	GetCode(error) int
}
