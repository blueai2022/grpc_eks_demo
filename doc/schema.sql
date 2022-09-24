-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-09-24T13:54:28.502Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "agency" varchar,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "clients" (
  "id" bigserial PRIMARY KEY,
  "agent_id" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "birth_date" timestamptz NOT NULL DEFAULT '0001-01-01',
  "driver_license_number" varchar,
  "driver_license_state" varchar,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint NOT NULL,
  "address_line1" varchar NOT NULL,
  "address_line2" varchar,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "zip_code" varchar NOT NULL,
  "country" varchar NOT NULL DEFAULT 'US',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "applications" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "primary_insured_id" bigint NOT NULL,
  "joint_insured_id" bigint,
  "carriers" varchar NOT NULL,
  "product_type_generic" varchar NOT NULL,
  "applied_amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "documents" (
  "id" bigserial PRIMARY KEY,
  "application_id" bigint NOT NULL,
  "is_primary_insured" boolean NOT NULL,
  "type" varchar NOT NULL,
  "file_size" varchar NOT NULL,
  "s3_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "applications" ("username");

CREATE INDEX ON "documents" ("application_id");

CREATE INDEX ON "documents" ("application_id", "type");

COMMENT ON COLUMN "applications"."product_type_generic" IS 'non carrier-specific product type';

COMMENT ON COLUMN "documents"."s3_url" IS 's3 bucket url';

ALTER TABLE "clients" ADD FOREIGN KEY ("agent_id") REFERENCES "users" ("username");

ALTER TABLE "addresses" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "applications" ADD FOREIGN KEY ("primary_insured_id") REFERENCES "clients" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("joint_insured_id") REFERENCES "clients" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");
