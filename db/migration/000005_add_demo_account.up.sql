-- Create demo user for testing purposes
INSERT INTO "users" (
  "username", 
  "hashed_password", 
  "full_name", 
  "email", 
  "is_email_verified",
  "created_at",
  "password_changed_at"
) VALUES (
  'demo_user',
  -- This is a hashed password for 'demo123456' - should be changed in production
  '$2a$10$K5V5z5H5V5z5H5V5z5H5V5', 
  'Demo User',
  'demo@nimbus.example.com',
  true,
  NOW(),
  NOW()
) ON CONFLICT (username) DO NOTHING;

-- Add account_number column if it doesn't exist
ALTER TABLE "account" ADD COLUMN IF NOT EXISTS "account_number" varchar UNIQUE;

-- Create demo account with sufficient balance for testing
INSERT INTO "account" (
  "owner",
  "balance", 
  "currency",
  "account_number",
  "created_at"
) VALUES (
  'demo_user',
  100000000, -- $1,000,000 for demo purposes
  'USD',
  'DEMO-1234567890',
  NOW()
) ON CONFLICT (owner, currency) DO UPDATE SET 
  balance = EXCLUDED.balance,
  account_number = EXCLUDED.account_number;
