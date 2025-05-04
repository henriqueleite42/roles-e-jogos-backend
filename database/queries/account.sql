-- name: UpdateAccountName :exec
UPDATE "accounts" SET
	"name" = $2
WHERE
	"id" = $1;

-- name: UpdateAccountHandle :exec
UPDATE "accounts" SET
	"handle" = $2
WHERE
	"id" = $1;

-- name: UpdateAccountProfile :exec
UPDATE "accounts" SET
	"name" = COALESCE($2, "name")
WHERE "id" = $1;

-- name: GetAccountById :one
SELECT
	a.*
FROM "accounts" a
WHERE
	a."id" = $1
LIMIT 1;

-- name: GetAccountByHandle :one
SELECT
	a.*
FROM "accounts" a
WHERE
	a."handle" = $1
LIMIT 1;

-- name: GetAccountsListByIds :many
SELECT
	a."avatar_path",
	a."created_at",
	a."handle",
	a."id",
	a."name"
FROM "accounts" a
WHERE
	a."id" = ANY($1::int[])
LIMIT 1;

-- name: GetAccountDataById :one
SELECT
	a."id",
	a."is_admin"
FROM "accounts" a
WHERE
	a."id" = $1
LIMIT 1;

-- name: GetAccountDataByHandle :one
SELECT
	a."id",
	a."is_admin"
FROM "accounts" a
WHERE
	a."handle" = $1
LIMIT 1;

-- name: GetAccountDataByEmail :one
SELECT
	e."account_id",
	a."is_admin"
FROM
	"email_addresses" e
LEFT JOIN "accounts" a ON e."account_id" = a."id"
WHERE
	e."email_address" = $1
LIMIT 1;

-- name: GetAccountDataByConnection :one
SELECT
	c."account_id",
	a."is_admin"
FROM "connections" c
LEFT JOIN "accounts" a ON c."account_id" = a."id"
WHERE
	c."provider" = $1
	AND
	c."external_id" = $2
LIMIT 1;

-- name: GetAccountDataByEmailOrConnection :one
(SELECT
	c."account_id",
	a."is_admin"
FROM
	"connections" c
LEFT JOIN "accounts" a ON c."account_id" = a."id"
WHERE
	c."provider" = $1
	AND
	c."external_id" = $2
LIMIT 1)
UNION
(SELECT
	e."account_id",
	a."is_admin"
FROM
	"email_addresses" e
LEFT JOIN "accounts" a ON e."account_id" = a."id"
WHERE
	e."email_address" = $3
LIMIT 1);

-- name: CreateAccountWithName :one
INSERT INTO "accounts" (
	"handle",
	"name"
) VALUES (
	$1,
	$2
) RETURNING "id";

-- name: CreateValidatedEmailAddress :exec
INSERT INTO "email_addresses" (
	"account_id",
	"email_address",
	"validated_at"
) VALUES (
	$1,
	$2,
	NOW()
) ON CONFLICT DO NOTHING;

-- name: GetEmailsByAccountsIds :many
SELECT DISTINCT ON (ea."account_id")
	ea."email_address",
	ea."account_id",
	ea."created_at",
	ea."validated_at"
FROM "email_addresses" ea
WHERE
	ea."account_id" = ANY($1::int[])
ORDER BY
	ea."created_at" ASC;

-- name: GetValidatedEmailsByAccountsIds :many
SELECT DISTINCT ON (ea."account_id")
	ea."email_address",
	ea."account_id",
	ea."created_at",
	ea."validated_at"
FROM "email_addresses" ea
WHERE
	ea."account_id" = ANY($1::int[])
	AND
	ea."validated_at" IS NOT NULL
ORDER BY
	ea."created_at" ASC;

-- name: CreateConnection :exec
INSERT INTO "connections" (
	"account_id",
	"external_handle",
	"external_id",
	"provider",
	"refresh_token"
) VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
);

-- name: UpdateConnection :exec
UPDATE "connections" SET
	"refresh_token" = COALESCE($4, "refresh_token"),
	"external_handle" = COALESCE($5, "external_handle")
WHERE
	"account_id" = $1
	AND
	"provider" = $2
	AND
	"external_id" = $3;

-- name: CreateOtp :exec
INSERT INTO "one_time_passwords" (
	"account_id",
	"code",
	"purpose"
) VALUES (
	$1,
	$2,
	$3
);

-- name: GetOtp :one
SELECT
	otp."created_at"
FROM "one_time_passwords" otp
WHERE
	otp."account_id" = $1
	AND
	otp."code" = $2
	AND
	otp."purpose" = $3
LIMIT 1;
