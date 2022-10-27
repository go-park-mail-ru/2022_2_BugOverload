--  PostgreSQL автоматически создает индексы по PRIMARY KEY
CREATE INDEX IF NOT EXISTS idx_film_name_prod_date_rating
    ON films USING btree
        (name, prod_date, rating);

CREATE INDEX IF NOT EXISTS idx_film_name
    ON films (LOWER(name));

CREATE INDEX IF NOT EXISTS idx_person_name
    ON persons (LOWER(name));

CREATE INDEX IF NOT EXISTS idx_collection_name_create_date_is_public
    ON collections USING btree
        (name, create_date, count_likes, is_public);

