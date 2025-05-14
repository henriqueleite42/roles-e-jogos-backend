-- name: CreateGame :one
INSERT INTO "games" (
	"name",
	"description",
	"icon_path",
	"kind",
	"ludopedia_id",
	"ludopedia_url",
	"min_amount_of_players",
	"max_amount_of_players",
	"average_duration",
	"min_age"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7,
	$8,
	$9,
	$10
) RETURNING "id";

-- name: GetGameByLudopediaId :one
SELECT
	g."id",
	g."name",
	g."description",
	g."icon_path",
	g."kind",
	g."ludopedia_id",
	g."ludopedia_url",
	g."min_amount_of_players",
	g."max_amount_of_players",
	g."average_duration",
	g."min_age",
	g."created_at"
FROM "games" g
WHERE
	g."ludopedia_id" = $1
LIMIT 1;

-- name: GetGamesListByLudopediaId :many
SELECT
	g."id",
	g."name",
	g."description",
	g."icon_path",
	g."kind",
	g."ludopedia_id",
	g."ludopedia_url",
	g."min_amount_of_players",
	g."max_amount_of_players",
	g."average_duration",
	g."min_age",
	g."created_at"
FROM "games" g
WHERE
	g."ludopedia_id" = ANY($1::int[]);
