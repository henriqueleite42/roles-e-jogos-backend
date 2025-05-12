-- name: GetCollectiveCollection :many
SELECT
  g."id" AS "game_id",
  g."name",
  g."description",
  g."kind",
  g."min_amount_of_players",
  g."max_amount_of_players",
  g."ludopedia_id",
  g."ludopedia_url",
  g."icon_path",
  json_agg(
    json_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path"
    ) ORDER BY a."handle"
  ) AS owners
FROM
  "games" g
INNER JOIN
  "personal_collections" pc ON g."id" = pc."game_id"
INNER JOIN
  "accounts" a ON pc."account_id" = a."id"
WHERE
	g."kind" = $1
	AND
	g."name" > $2
GROUP BY
  g."id"
ORDER BY
  g."name" ASC
LIMIT $3;

-- name: GetCollectiveCollectionByOwner :many
SELECT
  g."id" AS "game_id",
  g."name",
  g."description",
  g."kind",
  g."min_amount_of_players",
  g."max_amount_of_players",
  g."ludopedia_id",
  g."ludopedia_url",
  g."icon_path",
  json_agg(
    json_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path"
    ) ORDER BY a."handle"
  ) AS owners
FROM
  "games" g
INNER JOIN
  "personal_collections" pc ON g."id" = pc."game_id"
INNER JOIN
  "accounts" a ON pc."account_id" = a."id"
WHERE
	g."kind" = $1
	AND
	pc."account_id" = $2
	AND
	g."name" > $3
GROUP BY
  g."id"
ORDER BY
	g."name" ASC
LIMIT $4;

-- name: GetCollectiveCollectionByMaxAmountOfPlayers :many
SELECT
  g."id" AS "game_id",
  g."name",
  g."description",
  g."kind",
  g."min_amount_of_players",
  g."max_amount_of_players",
  g."ludopedia_id",
  g."ludopedia_url",
  g."icon_path",
  json_agg(
    json_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path"
    ) ORDER BY a."handle"
  ) AS owners
FROM
  "games" g
INNER JOIN
  "personal_collections" pc ON g."id" = pc."game_id"
INNER JOIN
  "accounts" a ON pc."account_id" = a."id"
WHERE
	g."kind" = $1
	AND
	g."max_amount_of_players" >= $2
	AND
	g."name" > $3
GROUP BY
  g."id"
ORDER BY
	g."name" ASC
LIMIT $4;

-- name: GetCollectiveCollectionByGameName :many
SELECT
  g."id" AS "game_id",
  g."name",
  g."description",
  g."kind",
  g."min_amount_of_players",
  g."max_amount_of_players",
  g."ludopedia_id",
  g."ludopedia_url",
  g."icon_path",
  json_agg(
    json_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path"
    ) ORDER BY a."handle"
  ) AS owners
FROM
  "games" g
INNER JOIN
  "personal_collections" pc ON g."id" = pc."game_id"
INNER JOIN
  "accounts" a ON pc."account_id" = a."id"
WHERE
	g."kind" = $1
	AND
	g."name" LIKE $2
	AND
	g."name" > $3
GROUP BY
  g."id"
ORDER BY
	g."name" ASC
LIMIT $4;

-- name: GetCollectiveCollectionAllFilters :many
SELECT
  g."id" AS "game_id",
  g."name",
  g."description",
  g."kind",
  g."min_amount_of_players",
  g."max_amount_of_players",
  g."ludopedia_id",
  g."ludopedia_url",
  g."icon_path",
  json_agg(
    json_build_object(
      'account_id', a."id",
      'handle', a."handle",
      'avatar_path', a."avatar_path"
    ) ORDER BY a."handle"
  ) AS owners
FROM
  "games" g
INNER JOIN
  "personal_collections" pc ON g."id" = pc."game_id"
INNER JOIN
  "accounts" a ON pc."account_id" = a."id"
WHERE
	g."kind" = $1
	AND
	g."name" LIKE $2
	AND
	pc."account_id" = $3
	AND
	g."max_amount_of_players" >= $4
	AND
	g."name" > $5
GROUP BY
  g."id"
ORDER BY
	g."name" ASC
LIMIT $6;
