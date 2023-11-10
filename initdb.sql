CREATE DATABASE cs_old;
CREATE DATABASE cs_new;

\c cs_old;

CREATE TABLE "public"."user" (
  "id" serial,
  "created_at" int,
  "updated_at" int,
  "deleted_at" int,
  "email_address" text,
  "password" text,
  "salt" text,
  "first_name" text,
  "last_name" text,
  "is_active" int,
  CONSTRAINT "user_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "user" ("id", "created_at", "updated_at", "deleted_at", "email_address", "password", "salt", "first_name", "last_name", "is_active") VALUES
(1,	1699619627,	1699619100,	NULL,	'john.smith@commonsense-pro.com',	'8933e4f0de98fe7cde084dfc171b6dacb1cc8e39',	'0948nyt493',	'John',	'Smith',	1),
(2,	1699624861,	1699618700,	NULL,	'richard.grant@commonsense-pro.com',	'f4b71f5480185e4c68fb37bb98ae2e87f6fcc946',	'gh984m984b',	'Richard',	'Grant',	1),
(3,	1699414271,	1699616720,	NULL,	'adam.conan@commonsense-pro.com',	'9f901942701e9060669cfa4e51c026c0e3125867',	'mviroijrp2',	'Adam',	'Conan',	0),
(4,	1699184337,	1699185637,	NULL,	'bill.rothman@commonsense-pro.com',	'035da764b8723cb4b5809790fc61f1077df19cdc',	'h34930jg08',	'Bill',	'Rothman',	1),
(5,	1699184337,	1699185637,	NULL,	'johannes.cristensen@commonsense-pro.com', '34da8e9bacf21be1930cae69d529e21874847799', 'a201dcbb32', 'Johannes', 'Cristensen', 1);

\c cs_new;

CREATE TABLE "public"."user_accounts" (
  "id" text NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz,
  "updated_at" timestamptz,
  "deleted_at" timestamptz,
  "email" text,
  "password" text,
  "salt" text,
  "first_name" text,
  "last_name" text,
  "is_active" boolean,
  "old_id" bigint,
  CONSTRAINT "user_accounts_pkey" PRIMARY KEY ("id"),
  CONSTRAINT "idx_user_account_email" UNIQUE ("email")
) WITH (oids = false);