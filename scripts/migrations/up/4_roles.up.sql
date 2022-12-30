-- Microservice user (while in the api microservice) - all user activity in app
CREATE ROLE user_app WITH LOGIN PASSWORD 'pass_microservice_2022_user';

-- Access
GRANT
    SELECT ON users, user_collections, user_ratings, user_reviews,user_views_films,
    collections, collections_films, collection_likes
    TO user_app;

GRANT
    UPDATE ON users, user_collections, user_ratings, user_reviews,
    collections, reviews, films
    TO user_app;

GRANT
    INSERT, DELETE ON reviews, user_reviews, user_ratings, collections,
    collections_films, collection_likes
    TO user_app;

-- Microservice auth
CREATE ROLE auth_app WITH LOGIN PASSWORD 'pass_microservice_2022_auth';

-- Access
GRANT
    SELECT, INSERT, UPDATE ON users TO auth_app;

-- For create default collections (at the moment)
GRANT
    INSERT ON collections, user_collections TO auth_app;

-- Microservice warehouse - all public content in app
CREATE ROLE warehouse_app WITH LOGIN PASSWORD 'pass_microservice_2022_warehouse';

-- Access
GRANT
    SELECT ON
    -- For film
    films, serials, reviews, user_reviews, users, genres, companies, countries,
    tags, persons, film_tags, film_images, film_persons, media, film_companies, film_countries, film_genres,
    -- For collections
    collections, user_collections, users,
    -- For persons
    persons, professions, person_professions, person_genres
    TO warehouse_app;


