-- Microservice user (while in the api microservice) - all user activity in app
-- Access
GRANT
    SELECT ON users, user_collections, user_ratings, user_reviews, user_views_films,
    collections, collections_films, collection_likes, collections_genres, reviews,
    films, media
    TO user_app;

GRANT
    UPDATE ON users, user_ratings, collections, reviews, films
    TO user_app;

GRANT
    INSERT, DELETE, UPDATE ON reviews, user_reviews, user_ratings, collections,
    collections_films, collection_likes, user_collections, users
    TO user_app;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO user_app;

-- Microservice auth
-- Access
GRANT
    SELECT, INSERT, UPDATE ON users TO auth_app;

-- For create default collections (at the moment)
GRANT
    SELECT, INSERT ON collections, user_collections TO auth_app;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO auth_app;

-- Microservice warehouse - all public content in app
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

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO warehouse_app;

-- Microservice image - for control image workflow
-- Access
GRANT
    UPDATE,SELECT (avatar, user_id) ON users TO image_app;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO image_app;



