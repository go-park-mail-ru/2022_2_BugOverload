package httpwrapper

// ErrResponse is structure for giving answers with errors.
type ErrResponse struct {
	ErrMassage string `json:"error,omitempty" example:"{{Area}}: [{{Reason}}]"`
}
