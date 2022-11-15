--  PostgreSQL автоматически создает индексы по PRIMARY KEY

CREATE INDEX IF NOT EXISTS idx_film_name_prod_date_rating
    ON films (rating, name, prod_year);

CREATE INDEX IF NOT EXISTS idx_users_email
    ON users (email);

CREATE INDEX IF NOT EXISTS idx_tags_name
    ON tags (name);

CREATE INDEX IF NOT EXISTS idx_film_name
    ON films (LOWER(name));

CREATE INDEX IF NOT EXISTS idx_person_name
    ON persons (LOWER(name));
