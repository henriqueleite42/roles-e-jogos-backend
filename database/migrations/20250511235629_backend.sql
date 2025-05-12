-- Modify "games" table
ALTER TABLE "games" ALTER COLUMN "icon_path" TYPE character varying(256), ALTER COLUMN "ludopedia_url" TYPE character varying(512), ALTER COLUMN "name" TYPE character varying(128);
