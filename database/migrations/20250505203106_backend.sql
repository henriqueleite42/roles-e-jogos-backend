-- Create "sessions" table
CREATE TABLE "sessions" ("account_id" integer NOT NULL, "created_at" timestamptz NOT NULL DEFAULT now(), "session_id" character(32) NOT NULL, PRIMARY KEY ("session_id"), CONSTRAINT "sessions_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Drop enum type "event_confirmation_status_enum"
DROP TYPE "event_confirmation_status_enum";
