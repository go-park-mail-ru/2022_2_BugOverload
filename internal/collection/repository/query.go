package repository

const (
	getFilmsByTag = `
SELECT f.film_id,
       f.name,
       f.original_name,
       f.prod_date,
       f.poster_ver,
       f.type,
       f.rating
FROM films f
         JOIN film_tags ft on f.film_id = ft.fk_film_id
         JOIN tags t on ft.fk_tag_id = t.tag_id
WHERE t.name = $1 AND f.rating < $2
GROUP BY f.film_id, f.rating
ORDER BY f.rating DESC
LIMIT $3`
)
