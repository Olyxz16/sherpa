-- Create "User" table
CREATE TABLE "User" ("id" serial NOT NULL, "username" text NOT NULL, "masterkey" text NULL DEFAULT '', "b64salt" text NULL DEFAULT '', "b64filekey" text NULL DEFAULT '', PRIMARY KEY ("id"));
-- Create "Auth" table
CREATE TABLE "Auth" ("id" serial NOT NULL, "user_id" integer NULL, "source" character varying(255) NULL, "access_token" character varying(255) NULL, "expires_in" double precision NULL, "refresh_token" character varying(255) NULL, "rt_expires_in" double precision NULL, PRIMARY KEY ("id"), CONSTRAINT "Auth_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "User" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "File" table
CREATE TABLE "File" ("owner_id" integer NOT NULL, "source" text NOT NULL, "reponame" text NOT NULL, "filename" text NOT NULL, "b64content" text NULL, "b64nonce" text NULL, PRIMARY KEY ("owner_id", "source", "reponame", "filename"), CONSTRAINT "File_owner_id_fkey" FOREIGN KEY ("owner_id") REFERENCES "User" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
