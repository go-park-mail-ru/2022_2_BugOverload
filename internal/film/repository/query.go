package repository

const (
	getFilmByID = `
SELECT name,
       original_name,
       prod_year,
       slogan,
       description,
       age_limit,
       duration,
       poster_ver,
       budget,
       box_office,
       currency_budget,
       count_seasons,
       end_year,
       type,
       rating,
       count_actors,
       count_scores,
       count_negative_reviews,
       count_neutral_reviews,
       count_positive_reviews
FROM films
WHERE film_id = $1`

	getFilmGenres = `
SELECT g.name
FROM genres g
         JOIN film_genres fg on g.genre_id = fg.fk_genre_id
WHERE fg.fk_film_id = $1
ORDER BY weight DESC`

	getFilmCompanies = `
SELECT c.name
FROM companies c
         JOIN film_companies fc on c.company_id = fc.fk_company_id
WHERE fc.fk_film_id = $1
ORDER BY weight DESC`

	getFilmCountries = `
SELECT c.name
FROM countries c
         JOIN film_countries fc on c.country_id = fc.fk_country_id
WHERE fc.fk_film_id = $1
ORDER BY weight DESC`

	getFilmTags = `
SELECT t.name
FROM tags t
         JOIN film_tags ft on t.tag_id = ft.fk_tag_id
WHERE ft.fk_film_id = $1
ORDER BY t.name DESC`

	getFilmImages = `
SELECT images_list
FROM film_images
WHERE film_id = $1`

	getFilmActors = `
SELECT fp.fk_person_id, p.name, p.avatar, fp.character
FROM persons p
         JOIN film_persons fp on p.person_id = fp.fk_person_id
WHERE fp.fk_film_id = $1
  AND fp.character IS NOT NULL
GROUP BY fp.fk_film_id, fp.weight, fp.character, p.name, fp.fk_person_id, p.avatar
ORDER BY fp.weight DESC`

	getFilmPersons = `
SELECT fp.fk_person_id, p.name, fp.fk_profession_id
FROM persons p
         JOIN film_persons fp on p.person_id = fp.fk_person_id
WHERE fp.fk_film_id = $1
  AND fp.character IS NULL
GROUP BY fp.fk_film_id, fp.fk_profession_id, fp.weight, p.name, fp.fk_person_id
ORDER BY fp.fk_profession_id, fp.weight DESC`
)