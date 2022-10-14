package httpmodels

type ErrResponseAuthNoCookie struct {
	ErrMassage string `json:"error,omitempty" example:"Auth: [request has no cookies]"`
}

type ErrResponseAuthNoSuchCookie struct {
	ErrMassage string `json:"error,omitempty" example:"Auth: [no such cookie]"`
}

type ErrResponseAuthDefault struct {
	ErrMassage string `json:"error,omitempty" example:"Auth: [{{Reason}}]"`
}

type ErrResponseAuthNoSuchUser struct {
	ErrMassage string `json:"error,omitempty" example:"Auth: [such user doesn't exist]"`
}
