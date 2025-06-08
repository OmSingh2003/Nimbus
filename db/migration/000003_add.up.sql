CREATE TABLE IF NOT EXISTS "sessions" (
  "id" uuid  PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar  NOT NULL,
  "is_boolean" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

-- Add foreign key constraint only if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint 
        WHERE conname = 'account_owner_fkey'
    ) THEN
        ALTER TABLE "account" ADD CONSTRAINT "account_owner_fkey" FOREIGN KEY ("owner") REFERENCES "users" ("username");
    END IF;
END$$;

-- Add unique constraint only if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint 
        WHERE conname = 'owner_currency_key'
    ) THEN
        ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");
    END IF;
END$$;
