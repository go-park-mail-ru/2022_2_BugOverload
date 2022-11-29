package search

const (
	searchFilmsByName = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
WHERE (f.name ILIKE $1 OR f.original_name ILIKE $1) AND f.type = 'film'
LIMIT 6`

	searchSeriesByName = `
SELECT f.film_id,
       f.name,
       f.original_name,
       extract(YEAR FROM f.prod_date),
       f.poster_ver,
       f.type,
       f.rating
FROM films f
WHERE (f.name ILIKE $1 OR f.original_name ILIKE $1) AND f.type = 'serial'
LIMIT 6`

	searchPersonsByName = `
SELECT
    person_id,
    name,
    birthday,
    original_name,
    avatar,
    count_films
FROM persons
WHERE name ILIKE $1 OR original_name ILIKE $1
LIMIT 6`

	getPersonProfessions = `
SELECT p.name
FROM professions p
         JOIN person_professions pp on p.profession_id = pp.fk_profession_id
WHERE pp.fk_person_id = $1
ORDER BY pp.weight DESC`
)
