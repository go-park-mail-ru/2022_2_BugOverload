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
	ErrMassage string `json:"error,omitempty" example:"Auth: [no such user]"`
}

type ErrResponseAuthWrongLoginCombination struct {
	ErrMassage string `json:"error,omitempty" example:"Auth: [no such combination of login and password]"`
}

type ErrResponseImageNoSuchImage struct {
	ErrMassage string `json:"error,omitempty" example:"Image: [no such image]"`
}

type ErrResponseImageDefault struct {
	ErrMassage string `json:"error,omitempty" example:"Image: [{{Reason}}]"`
}

type ErrResponseFilmNoSuchFilm struct {
	ErrMassage string `json:"error,omitempty" example:"Film: [no such film]"`
}

type ErrResponsePersonNoSuchPerson struct {
	ErrMassage string `json:"error,omitempty" example:"Person: [no such person]"`
}

type ErrResponseCollectionNoSuchCollection struct {
	ErrMassage string `json:"error,omitempty" example:"Collection: [no such collection]"`
}

type ErrResponseProfileNoSuchProfile struct {
	ErrMassage string `json:"error,omitempty" example:"Profile: [no such profile]"`
}
