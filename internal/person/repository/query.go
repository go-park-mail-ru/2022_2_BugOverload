package repository

const (
	getPerson = `SELECT 
    name, 
    birthday, 
    growth, 
    original_name, 
    avatar, 
    death, 
    gender, 
    count_films  
	FROM persons
	WHERE person_id = $1`

	getPersonBestFilms = `SELECT 
    	f.film_id, 
    	f.name, 
    	f.original_name, 
        f.prod_year,  
        f.poster_ver, 
        f.end_year, 
        f.rating
		FROM
    		films f
       		JOIN film_persons fp ON f.film_id = fp.fk_film_id
        	JOIN persons p ON fp.fk_person_id = $1
		GROUP BY f.film_id
		LIMIT $2`

	getPersonImages = `SELECT images_list  FROM person_images WHERE person_id = $1`
)
