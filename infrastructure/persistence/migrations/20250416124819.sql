-- Modify "Auth" table
ALTER TABLE "public"."Auth" DROP CONSTRAINT "Auth_pkey", ALTER COLUMN "user_id" SET NOT NULL, ADD PRIMARY KEY ("id", "user_id");
-- Create index "index_userid_source" to table: "Auth"
CREATE UNIQUE INDEX "index_userid_source" ON "public"."Auth" ("user_id", "source");
