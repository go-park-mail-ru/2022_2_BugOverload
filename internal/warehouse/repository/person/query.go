package person

const (
	GetPersonByID = `SELECT 
    name, 
    birthday, 
    growth_meters, 
    original_name, 
    avatar, 
    death, 
    gender, 
    count_films  
	FROM persons
	WHERE person_id = $1`

	GetPersonBestFilms = `SELECT 
    	f.film_id, 
    	f.name, 
    	f.original_name, 
        extract(YEAR FROM  f.prod_date),
        f.poster_ver, 
        f.type, 
        f.rating
		FROM
    		films f
       		JOIN film_persons fp ON f.film_id = fp.fk_film_id
        	JOIN persons p ON fp.fk_person_id = $1
		GROUP BY f.film_id, f.rating
		ORDER BY f.rating DESC
		LIMIT $2`

	GetPersonImages = `SELECT image_key FROM person_images WHERE person_id = $1 ORDER BY weight DESC LIMIT $2`

	GetPersonProfessions = `
SELECT p.name
FROM professions p
         JOIN person_professions pp on p.profession_id = pp.fk_profession_id
WHERE pp.fk_person_id = $1
ORDER BY pp.weight DESC`

	GetPersonGenres = `
SELECT g.name
FROM genres g
         JOIN person_genres pg on g.genre_id = pg.fk_genre_id
WHERE pg.fk_person_id = $1
ORDER BY pg.weight DESC`
)
