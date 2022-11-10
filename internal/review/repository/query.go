package repository

const (
	getReviewsByFilmID = `
SELECT r.name,
       r.type,
       r.body,
       r.count_likes,
       r.create_time,
       u.user_id,
       u.nickname,
       p.avatar,
       p.count_reviews
FROM reviews r
         JOIN profile_reviews pr on r.review_id = pr.fk_review_id
         JOIN profiles p on pr.fk_profile_id = p.profile_id
         JOIN users u on p.profile_id = u.user_id
WHERE pr.fk_film_id = $1
ORDER BY r.count_likes IS NULL, r.count_likes DESC
LIMIT $2 OFFSET $3`
)
