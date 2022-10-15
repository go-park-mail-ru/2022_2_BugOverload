package tests

// TestCase is structure for API testing
type TestCase struct {
	Method          string
	ContentType     string
	RequestBody     string
	Cookie          string
	CookieUserEmail string
	ResponseCookie  string
	ResponseBody    string
	StatusCode      int
}
