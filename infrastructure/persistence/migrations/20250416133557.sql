-- Modify "Auth" table
ALTER TABLE "public"."Auth" DROP CONSTRAINT "Auth_pkey", ALTER COLUMN "source" SET NOT NULL, ADD PRIMARY KEY ("id", "user_id", "source");
