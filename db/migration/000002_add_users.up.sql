CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hash_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "emial_id" varchar UNIQUE NOT NULL,
  "password_last_changed" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
