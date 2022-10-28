package models

type Film struct {
	ID                   uint     `json:"id,omitempty"`
	Name                 string   `json:"name,omitempty"`
	ProdYear             int      `json:"prod_year,omitempty"`
	Type                 string   `json:"type,omitempty"`
	Description          string   `json:"description,omitempty"`
	ShortDescription     string   `json:"short_description,omitempty"`
	AgeLimit             string   `json:"age_limit,omitempty"`
	BoxOffice            int      `json:"box_office,omitempty"`
	Duration             string   `json:"duration,omitempty"`
	PosterHor            string   `json:"poster_hor,omitempty"`
	PosterVer            string   `json:"poster_ver,omitempty"`
	CountSeasons         int      `json:"count_seasons,omitempty"`
	EndYear              int      `json:"end_year,omitempty"`
	Rating               float32  `json:"rating,omitempty"`
	CountScores          int      `json:"count_scores,omitempty"`
	CountNegativeReviews int      `json:"count_negative_reviews,omitempty"`
	CountNeutralReviews  int      `json:"count_neutral_reviews,omitempty"`
	CountPositiveReviews int      `json:"count_positive_reviews,omitempty"`
	Genres               []string `json:"genres,omitempty"`
	ProdCompanies        []string `json:"prod_companies,omitempty"`
	ProdCountries        []string `json:"prod_countries,omitempty"`
}
