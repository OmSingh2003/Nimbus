-- Add is_email_verified column to users table
ALTER TABLE "users" ADD COLUMN "is_email_verified" bool NOT NULL DEFAULT false;

-- Create verify_emails table
CREATE TABLE "verify_emails" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT (now() + interval '24 hours')
);

-- Add foreign key constraint
ALTER TABLE "verify_emails" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
