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
  AND ($2 = '' OR LOWER(g."name") LIKE LOWER('%' || $2 || '%'))
  AND ($3 = 0 OR pc."account_id" = $3)
  AND ($4 = 0 OR g."max_amount_of_players" >= $4)
  AND ($5 = '' OR g."name" > $5)
GROUP BY
  g."id"
ORDER BY
	g."name" ASC
LIMIT $6;

-- name: AddToPersonalCollection :exec
INSERT INTO "personal_collections" (
	"account_id",
	"game_id",
	"paid",
	"acquired_at"
) VALUES (
	$1,
	$2,
	$3,
	$4
) ON CONFLICT ("account_id", "game_id") DO NOTHING;

-- name: GetOngoingImportCollectionLog :many
SELECT
	icl."id",
	icl."account_id",
	icl."external_id",
	icl."provider",
	icl."trigger",
	icl."status",
	icl."created_at"
FROM "import_collection_logs" icl
WHERE
	icl."external_id" = ANY($1::text[])
	AND icl."provider" = $2
	AND icl."ended_at" IS NULL;

-- name: CreateImportCollectionLog :one
INSERT INTO "import_collection_logs" (
	"account_id",
	"external_id",
	"trigger",
	"provider",
	"status"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
) RETURNING "id";

-- name: UpdateManyImportCollectionsLogs :exec
UPDATE "import_collection_logs"
SET
	"status" = $2
WHERE
	"external_id" = ANY($1::int[])
	AND "ended_at" IS NULL;
