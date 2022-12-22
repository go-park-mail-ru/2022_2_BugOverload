package wrapper

//go:generate easyjson  -disallow_unknown_fields models.go

// ErrResponse is structure for giving answers with errors.
//
//easyjson:json
type ErrResponse struct {
	ErrMassage string `json:"error,omitempty" example:"{{Area}}: [{{Reason}}]"`
}
