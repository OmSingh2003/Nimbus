-- Remove demo account
DELETE FROM "account" WHERE "account_number" = 'DEMO-1234567890';

-- Remove demo user
DELETE FROM "users" WHERE "username" = 'demo_user';

-- Remove account_number column (be careful in production)
-- ALTER TABLE "account" DROP COLUMN IF EXISTS "account_number";
