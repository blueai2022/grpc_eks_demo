CREATE TABLE "api_accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "is_active" boolean NOT NULL DEFAULT true,
  "is_auto_renewal" boolean NOT NULL DEFAULT false,
  "service_type" varchar NOT NULL DEFAULT 'ICD',
  "plan_name" varchar NOT NULL DEFAULT 'DEMO',
  "credit_balance" bigint NOT NULL,
  "active_at" timestamptz NOT NULL DEFAULT (now()),
  "last_use_at" timestamptz NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "api_accounts" ("username");

CREATE INDEX ON "api_accounts" ("username", "is_active");

CREATE UNIQUE INDEX ON "api_accounts" ("username", "is_active", "service_type");

COMMENT ON COLUMN "api_accounts"."service_type" IS 'ICD|ICD_PRO|APS|APS_TXT|ALL';

COMMENT ON COLUMN "api_accounts"."plan_name" IS 'DEMO|BASIC|PRO';

ALTER TABLE "api_accounts" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");