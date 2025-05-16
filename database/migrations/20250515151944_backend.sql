-- Add value to enum type: "collection_import_status_enum"
ALTER TYPE "collection_import_status_enum" ADD VALUE 'NOT_YET_STARTED';
-- Create enum type "location_kind_enum"
CREATE TYPE "location_kind_enum" AS ENUM ('BUSINESS', 'PERSONAL');
-- Modify "event_attendances" table
ALTER TABLE "event_attendances" DROP CONSTRAINT "event_attendances_pkey", ADD PRIMARY KEY ("id");
-- Create index "event_attendances_event_id_account_id_idx" to table: "event_attendances"
CREATE UNIQUE INDEX "event_attendances_event_id_account_id_idx" ON "event_attendances" ("event_id", "account_id");
-- Modify "event_games" table
ALTER TABLE "event_games" DROP CONSTRAINT "event_games_pkey", ADD COLUMN "owner_id" integer NOT NULL, ADD PRIMARY KEY ("id"), ADD CONSTRAINT "event_games_owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Create index "event_games_event_id_game_id_owner_id_idx" to table: "event_games"
CREATE UNIQUE INDEX "event_games_event_id_game_id_owner_id_idx" ON "event_games" ("event_id", "game_id", "owner_id");
-- Create "locations" table
CREATE TABLE "locations" ("address" character varying(500) NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "created_by" integer NOT NULL, "icon_path" character varying(256) NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "kind" "location_kind_enum" NOT NULL, "name" character varying(128) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "locations_created_by_fk" FOREIGN KEY ("created_by") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "locations_address_idx" to table: "locations"
CREATE INDEX "locations_address_idx" ON "locations" ("address");
-- Create index "locations_name_idx" to table: "locations"
CREATE INDEX "locations_name_idx" ON "locations" ("name");
-- Modify "events" table
ALTER TABLE "events" DROP COLUMN "location_address", DROP COLUMN "location_name", ADD COLUMN "end_date" timestamptz NULL, ADD COLUMN "location_id" integer NOT NULL, ADD CONSTRAINT "events_location_id_fk" FOREIGN KEY ("location_id") REFERENCES "locations" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Create index "events_name_idx" to table: "events"
CREATE INDEX "events_name_idx" ON "events" ("name");
-- Rename a column from "date" to "start_date"
ALTER TABLE "events" RENAME COLUMN "date" TO "start_date";
-- Create index "events_start_date_idx" to table: "events"
CREATE INDEX "events_start_date_idx" ON "events" ("start_date");
-- Rename a column from "max_amount_of_players" to "capacity"
ALTER TABLE "events" RENAME COLUMN "max_amount_of_players" TO "capacity";
