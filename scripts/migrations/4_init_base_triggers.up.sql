CREATE OR REPLACE FUNCTION create_default_user_env() RETURNS TRIGGER AS
$$
DECLARE
    user_id                    integer;
    collection_favorites       varchar(64) := 'Избранное';
    collection_favorites_desc  TEXT        := 'Ваши сохранные фильмы';
    collection_will_watch      varchar(64) := 'Буду смотреть';
    collection_will_watch_desc TEXT        := 'Фильмы, которые вы отметили для просмотра';
    id_col_1                   integer;
    id_col_2                   integer;
BEGIN
    user_id := NEW.user_id;
    INSERT INTO profiles(profile_id) VALUES (user_id);

    INSERT INTO collections(name, description)
    VALUES (collection_favorites, collection_favorites_desc)
    RETURNING collection_id INTO id_col_1;

    INSERT INTO collections(name, description)
    VALUES (collection_will_watch, collection_will_watch_desc)
    RETURNING collection_id INTO id_col_2;

    INSERT INTO profile_collections(fk_collection_id, fk_profile_id)
    VALUES (id_col_1, user_id),
           (id_col_2, user_id);

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER t_users_insert
    AFTER INSERT
    ON users
EXECUTE PROCEDURE create_default_user_env();
