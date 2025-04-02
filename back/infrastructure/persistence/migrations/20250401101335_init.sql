-- Create "userdata" table
CREATE TABLE "public"."userdata" ("uid" serial NOT NULL, "username" text NOT NULL, "masterkey" text NULL DEFAULT '', "b64salt" text NULL DEFAULT '', "b64filekey" text NULL DEFAULT '', PRIMARY KEY ("uid"));
-- Create "auth" table
CREATE TABLE "public"."auth" ("uid" serial NOT NULL, "userid" integer NOT NULL, "source" character varying(255) NULL, "access_token" character varying(255) NULL, "expires_in" double precision NULL, "refresh_token" character varying(255) NULL, "rt_expires_in" double precision NULL, PRIMARY KEY ("uid", "userid"), CONSTRAINT "auth_userid_fkey" FOREIGN KEY ("userid") REFERENCES "public"."userdata" ("uid") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "file" table
CREATE TABLE "public"."file" ("ownerid" integer NOT NULL, "source" text NOT NULL, "reponame" text NOT NULL, "filename" text NOT NULL, "b64content" text NULL, "b64nonce" text NULL, PRIMARY KEY ("ownerid", "source", "reponame", "filename"), CONSTRAINT "file_ownerid_fkey" FOREIGN KEY ("ownerid") REFERENCES "public"."userdata" ("uid") ON UPDATE NO ACTION ON DELETE NO ACTION);
