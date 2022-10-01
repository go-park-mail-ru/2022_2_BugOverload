package structs

// Film is a carrier structure for all movie attributes and specifying them for json conversion
type Film struct {
	ID          uint     `json:"film_id,omitempty"`
	Name        string   `json:"film_name,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"film_type,omitempty"`
	YearProd    string   `json:"year_prod,omitempty"`
	ProdCompany string   `json:"prod_company,omitempty"`
	ProdCountry string   `json:"prod_country,omitempty"`
	AgeLimit    string   `json:"age_limit,omitempty"`
	Duration    string   `json:"duration,omitempty"`
	PosterHor   string   `json:"poster_hor,omitempty"`
	PosterVer   string   `json:"poster_ver,omitempty"`
	Rating      string   `json:"film_rating,omitempty"`
	Genres      []string `json:"genres,omitempty"`
}
