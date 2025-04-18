-- Create the account table
CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  -- Corrected data type from 'timestamps' to 'timestamptz'
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create the entries table
CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  -- Corrected data type from 'timestamps' to 'timestamptz'
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create the transfers table
CREATE TABLE "transfers" (
  -- Changed primary key to bigserial for consistency, assuming IDs should auto-increment
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  -- Added constraint to ensure amount is positive, as per comment
  "amount" bigint NOT NULL CHECK (amount > 0),
  -- Corrected data type from 'timestamps' to 'timestamptz'
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create indexes
CREATE INDEX ON "account" ("owner");
CREATE INDEX ON "entries" ("account_id");
CREATE INDEX ON "transfers" ("from_account_id");
CREATE INDEX ON "transfers" ("to_account_id");
CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

-- Add comments
COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';
COMMENT ON COLUMN "transfers"."amount" IS 'It must be positive';

-- Add foreign key constraints
-- Ensure account_id in entries references a valid account
ALTER TABLE "entries" ADD CONSTRAINT fk_entries_account FOREIGN KEY ("account_id") REFERENCES "account" ("id");

-- Ensure from_account_id in transfers references a valid account
ALTER TABLE "transfers" ADD CONSTRAINT fk_transfers_from_account FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

-- Ensure to_account_id in transfers references a valid account
ALTER TABLE "transfers" ADD CONSTRAINT fk_transfers_to_account FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");

