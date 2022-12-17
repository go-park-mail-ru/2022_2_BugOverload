package repository

const (
	getRelease = `
SELECT f.film_id,
       f.name,
       to_char(f.prod_date, 'YYYY.MM.DD'),
       f.poster_ver,
       f.rating
FROM films f
WHERE f.prod_date > now()
ORDER BY f.prod_date, f.film_id
LIMIT 5`
)
