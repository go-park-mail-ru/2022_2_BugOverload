package repository

const (
	getFilmByTag = `SELECT f.film_id,
       f.name,
       f.original_name,
       f.prod_year,
       f.poster_ver,
       f.end_year,
       f.rating
FROM films f
         JOIN film_tags ft on f.film_id = ft.fk_film_id
         JOIN tags t on ft.fk_tag_id = t.tag_id
WHERE t.name = 'популярное' AND f.film_id > 11
GROUP BY f.film_id, f.rating
ORDER BY f.rating DESC
LIMIT 5`
)
