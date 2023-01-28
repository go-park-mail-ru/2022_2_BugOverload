-- Microservice user (while in the api microservice) - all user activity in app
-- Access
GRANT
    SELECT ON users, user_collections, user_ratings, user_reviews, user_views_films,
    collections, collections_films, collection_likes, collections_genres,
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
    -- Global
    genres, companies, countries, tags, professions,
    -- For reviews
    reviews, reviews_likes,
    -- For film
    films, serials, media, reviews, users, film_tags, film_images, film_persons, film_companies, film_countries, film_genres,
    -- For users
    users, user_collections, user_ratings, user_reviews, user_views_films,
    -- For collections
    collections, collection_likes, collections_films, collections_genres,
    -- For persons
    persons, person_professions, person_genres, person_images
    TO warehouse_app;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO warehouse_app;

-- Microservice image - for control image workflow
-- Access
GRANT
    UPDATE (avatar) ON users TO image_app;

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO image_app;



