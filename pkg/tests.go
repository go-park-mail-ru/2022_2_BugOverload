package pkg

import "fmt"

func TestErrorMessage(number int, message string) string {
	return fmt.Sprintf("Case[%d] FAIL: %s", number, message)
}
