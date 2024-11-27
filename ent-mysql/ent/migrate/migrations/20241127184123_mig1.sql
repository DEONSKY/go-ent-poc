-- Create "books" table
CREATE TABLE "books" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "delete_time" timestamptz NULL, "title" character varying NOT NULL, "author" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "delete_time" timestamptz NULL, "age" bigint NOT NULL, "rank" double precision NULL, "active" boolean NOT NULL DEFAULT false, "password" character varying NOT NULL, "name" character varying NOT NULL, "url" jsonb NULL, "strings" jsonb NULL, "state" bigint NOT NULL DEFAULT 0, "uuid" uuid NOT NULL, PRIMARY KEY ("id"));
-- Create index "user_age_name" to table: "users"
CREATE UNIQUE INDEX "user_age_name" ON "users" ("age", "name");
-- Create index "users_name_key" to table: "users"
CREATE UNIQUE INDEX "users_name_key" ON "users" ("name");
-- Create "cards" table
CREATE TABLE "cards" ("id" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "delete_time" timestamptz NULL, "card_no" character varying NOT NULL, "user_card" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "cards_users_card" FOREIGN KEY ("user_card") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
