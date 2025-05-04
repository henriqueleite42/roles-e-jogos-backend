-- Create enum type "event_confirmation_status_enum"
CREATE TYPE "event_confirmation_status_enum" AS ENUM ('GOING', 'MAYBE', 'NOT_GOING');
-- Create enum type "provider_enum"
CREATE TYPE "provider_enum" AS ENUM ('GOOGLE', 'LUDOPEDIA');
-- Create enum type "otp_purpose_enum"
CREATE TYPE "otp_purpose_enum" AS ENUM ('SIGN_IN');
-- Create enum type "kind_enum"
CREATE TYPE "kind_enum" AS ENUM ('RPG', 'GAME', 'EXPANSION');
-- Create enum type "event_attendance_status_enum"
CREATE TYPE "event_attendance_status_enum" AS ENUM ('GOING', 'MAYBE', 'NOT_GOING');
-- Create enum type "event_confidentiality_enum"
CREATE TYPE "event_confidentiality_enum" AS ENUM ('PUBLIC', 'ONLY_INVITED');
-- Create "accounts" table
CREATE TABLE "accounts" ("avatar_path" character(20) NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "handle" character varying(16) NOT NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "is_admin" boolean NOT NULL DEFAULT false, "name" character varying(50) NULL, PRIMARY KEY ("id"));
-- Create index "accounts_handle_idx" to table: "accounts"
CREATE UNIQUE INDEX "accounts_handle_idx" ON "accounts" ("handle");
-- Create "connections" table
CREATE TABLE "connections" ("account_id" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "external_handle" character varying(30) NULL, "external_id" character varying(255) NOT NULL, "provider" "provider_enum" NOT NULL, "refresh_token" character varying(500) NULL, PRIMARY KEY ("external_id", "provider"), CONSTRAINT "connections_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create index "connections_account_id_idx" to table: "connections"
CREATE INDEX "connections_account_id_idx" ON "connections" ("account_id");
-- Create "email_addresses" table
CREATE TABLE "email_addresses" ("account_id" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "email_address" character varying(500) NOT NULL, "validated_at" timestamptz NULL, PRIMARY KEY ("email_address"), CONSTRAINT "email_addresses_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "events" table
CREATE TABLE "events" ("created_at" timestamptz NOT NULL DEFAULT now(), "date" timestamptz NOT NULL, "description" character varying(1000) NOT NULL, "icon_path" character varying(250) NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "location_address" character varying(500) NOT NULL, "location_name" character varying(100) NOT NULL, "max_amount_of_players" integer NULL, "name" character varying(50) NOT NULL, "owner_id" integer NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "events_owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "event_attendances" table
CREATE TABLE "event_attendances" ("account_id" integer NOT NULL, "confirmed_at" timestamptz NOT NULL DEFAULT now(), "event_id" integer NOT NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "status" "event_attendance_status_enum" NOT NULL, PRIMARY KEY ("account_id", "event_id", "id"), CONSTRAINT "event_attendances_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "event_attendances_event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "games" table
CREATE TABLE "games" ("created_at" timestamptz NOT NULL DEFAULT now(), "description" character varying(1000) NOT NULL, "icon_path" character varying(250) NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "kind" "kind_enum" NOT NULL, "ludopedia_id" integer NULL, "ludopedia_url" character varying(500) NULL, "max_amount_of_players" integer NOT NULL, "min_amount_of_players" integer NOT NULL, "name" character varying(50) NOT NULL, PRIMARY KEY ("id"));
-- Create index "games_ludopedia_id_idx" to table: "games"
CREATE INDEX "games_ludopedia_id_idx" ON "games" ("ludopedia_id");
-- Create "event_games" table
CREATE TABLE "event_games" ("event_id" integer NOT NULL, "game_id" integer NOT NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, PRIMARY KEY ("event_id", "game_id", "id"), CONSTRAINT "event_games_event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "event_games_game_id_fk" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "medias" table
CREATE TABLE "medias" ("created_at" timestamptz NOT NULL DEFAULT now(), "date" timestamptz NOT NULL, "description" character varying(500) NULL, "game_id" integer NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "owner_id" integer NOT NULL, "path" character varying(500) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "medias_owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "one_time_passwords" table
CREATE TABLE "one_time_passwords" ("account_id" integer NOT NULL, "code" character varying(255) NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "purpose" "otp_purpose_enum" NOT NULL, PRIMARY KEY ("account_id", "code", "purpose"), CONSTRAINT "one_time_passwords_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "personal_collections" table
CREATE TABLE "personal_collections" ("account_id" integer NOT NULL, "acquired_at" timestamptz NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "game_id" integer NOT NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "paid" integer NULL, PRIMARY KEY ("id"), CONSTRAINT "personal_collections_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "personal_collections_game_id_fk" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
