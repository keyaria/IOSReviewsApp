CREATE TABLE "review" (
  "id" bigserial PRIMARY KEY,
  "uri" string NOT NULL,
  "name" string NOT NULL,
  "updated" timestamp NOT NULL,
  "rating" string NOT NULL,
  "title" string NOT NULL,
  "content" string NOT NULL
);

CREATE INDEX ON "review" ("updated");

CREATE INDEX ON "review" ("rating");

COMMENT ON COLUMN "review"."rating" IS 'Can not be negative';
