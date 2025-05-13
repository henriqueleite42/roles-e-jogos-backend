-- Modify "events" table
ALTER TABLE "events" DROP CONSTRAINT "events_owner_id_fk", ADD CONSTRAINT "events_owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "event_attendances" table
ALTER TABLE "event_attendances" DROP CONSTRAINT "event_attendances_account_id_fk", DROP CONSTRAINT "event_attendances_event_id_fk", ADD CONSTRAINT "event_attendances_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "event_attendances_event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "event_games" table
ALTER TABLE "event_games" DROP CONSTRAINT "event_games_event_id_fk", DROP CONSTRAINT "event_games_game_id_fk", ADD CONSTRAINT "event_games_event_id_fk" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "event_games_game_id_fk" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "medias" table
ALTER TABLE "medias" DROP CONSTRAINT "medias_owner_id_fk", ADD CONSTRAINT "medias_owner_id_fk" FOREIGN KEY ("owner_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "personal_collections" table
ALTER TABLE "personal_collections" DROP CONSTRAINT "personal_collections_account_id_fk", DROP CONSTRAINT "personal_collections_game_id_fk", ADD CONSTRAINT "personal_collections_account_id_fk" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "personal_collections_game_id_fk" FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
