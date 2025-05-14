-- Create enum type "collection_import_trigger_enum"
CREATE TYPE "collection_import_trigger_enum" AS ENUM ('ACCOUNT_CREATION', 'MANUAL_BY_USER');
-- Create enum type "collection_import_status_enum"
CREATE TYPE "collection_import_status_enum" AS ENUM ('STARTED', 'COMPLETED', 'FAILED');
-- Create index "personal_collections_account_id_game_id_idx" to table: "personal_collections"
CREATE UNIQUE INDEX "personal_collections_account_id_game_id_idx" ON "personal_collections" ("account_id", "game_id");
-- Create "import_collection_logs" table
CREATE TABLE "import_collection_logs" ("account_id" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "ended_at" timestamptz NULL, "external_id" character varying(255) NOT NULL, "id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "provider" "provider_enum" NOT NULL, "status" "collection_import_status_enum" NOT NULL, "trigger" "collection_import_trigger_enum" NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "import_collection_logs_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
