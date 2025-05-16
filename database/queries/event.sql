-- name: CreateEvent :one
INSERT INTO "events" (
	"owner_id",
	"name",
	"description",
	"icon_path",
	"start_date",
	"end_date",
	"capacity"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6,
	$7
) RETURNING "id";

-- name: CreateEventGame :exec
INSERT INTO "event_games" (
	"event_id",
	"game_id",
	"owner_id"
) VALUES (
	$1,
	$2,
	$3
);

-- name: CreateEventAttendance :exec
INSERT INTO "event_attendances" (
	"event_id",
	"account_id",
	"status"
) VALUES (
	$1,
	$2,
	$3
)
ON CONFLICT ("account_id", "event_id") DO UPDATE
SET "status" = $3;

-- name: GetNextEvents :many
SELECT
	e."id",
	e."owner_id",
	e."name",
	e."description",
	e."icon_path",
	e."start_date",
	e."end_date",
	e."capacity",
	l."id" AS "location_id",
	l."name" AS "location_name",
	l."address" AS "location_address",
	l."icon_path" AS "location_icon_path",
  json_agg(
    DISTINCT jsonb_build_object(
      'id', g."id",
      'name', g."name",
      'icon_path', g."icon_path",
      'kind', g."kind",
      'ludopedia_url', g."ludopedia_url",
      'min_amount_of_players', g."min_amount_of_players",
      'max_amount_of_players', g."max_amount_of_players",
      'average_duration', g."average_duration",
      'min_age', g."min_age"
    )
  ) AS "games",
  json_agg(
    DISTINCT jsonb_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path",
      'status', ea."status"
    )
  ) AS "attendances"
FROM "events" e
INNER JOIN "locations" l ON e."location_id" = l."id"
LEFT JOIN "event_games" eg ON eg."event_id" = e."id"
LEFT JOIN "games" g ON eg."game_id" = g."id"
LEFT JOIN "event_attendances" ea ON ea."event_id" = e."id"
LEFT JOIN "accounts" a ON ea."account_id" = a."id"
WHERE
	(
		e."start_date" >= $1
		OR
		e."end_date" >= $1
	)
GROUP BY
  e."id", l."id"
ORDER BY
	e."start_date" ASC
LIMIT $2;
