package structs

type Film struct {
	ID          uint   `json:"film_id,integer,omitempty"`
	Name        string `json:"film_name,string,omitempty"`
	Type        string `json:"film_type,string,omitempty"`
	YearProd    string `json:"year_prod,string,omitempty"`
	ProdCompany string `json:"prod_company,string,omitempty"`
	ProdCountry string `json:"prod_country,string,omitempty"`
	AgeLimit    string `json:"age_limit,string,omitempty"`
	Duration    string `json:"duration,string,omitempty"`
	PosterHor   string `json:"poster_hor,string,omitempty"`
	PosterVer   string `json:"poster_ver,string,omitempty"`
}
