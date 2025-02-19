-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-09-25T12:05:04.803Z

CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "address_id" bigint,
  "agency" varchar,
  "app_contact" varchar,
  "app_contact_email" varchar,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "clients" (
  "id" bigserial PRIMARY KEY,
  "agent" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "address_id" bigint,
  "birth_date" timestamptz NOT NULL,
  "driver_license_number" varchar,
  "driver_license_state" varchar,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
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
  "agent" varchar NOT NULL,
  "is_formal" boolean NOT NULL,
  "is_1035_exchange" boolean,
  "primary_insured_id" bigint NOT NULL,
  "joint_insured_id" bigint,
  "carriers" varchar NOT NULL,
  "product_type" varchar NOT NULL,
  "applied_amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "documents" (
  "id" bigserial PRIMARY KEY,
  "application_id" bigint NOT NULL,
  "client_id" bigint NOT NULL,
  "doc_type" varchar NOT NULL,
  "file_name" varchar NOT NULL,
  "file_size" varchar NOT NULL,
  "file_type" varchar NOT NULL,
  "s3_url" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "clients" ("agent");

CREATE INDEX ON "clients" ("agent", "full_name");

CREATE INDEX ON "clients" ("agent", "birth_date");

CREATE INDEX ON "clients" ("agent", "driver_license_number");

CREATE INDEX ON "addresses" ("address_line1");

CREATE INDEX ON "addresses" ("zip_code");

CREATE INDEX ON "applications" ("agent");

CREATE INDEX ON "applications" ("agent", "primary_insured_id");

CREATE INDEX ON "applications" ("agent", "joint_insured_id");

CREATE INDEX ON "documents" ("application_id");

CREATE INDEX ON "documents" ("application_id", "file_name");

COMMENT ON COLUMN "applications"."product_type" IS 'non carrier-specific product type';

COMMENT ON COLUMN "documents"."s3_url" IS 's3 bucket url';

ALTER TABLE "users" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "clients" ADD FOREIGN KEY ("agent") REFERENCES "users" ("username");

ALTER TABLE "clients" ADD FOREIGN KEY ("address_id") REFERENCES "addresses" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("agent") REFERENCES "users" ("username");

ALTER TABLE "applications" ADD FOREIGN KEY ("primary_insured_id") REFERENCES "clients" ("id");

ALTER TABLE "applications" ADD FOREIGN KEY ("joint_insured_id") REFERENCES "clients" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("application_id") REFERENCES "applications" ("id");

ALTER TABLE "documents" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");
