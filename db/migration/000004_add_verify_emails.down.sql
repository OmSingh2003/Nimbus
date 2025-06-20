-- Drop verify_emails table
DROP TABLE IF EXISTS "verify_emails";

-- Remove is_email_verified column from users table
ALTER TABLE "users" DROP COLUMN IF EXISTS "is_email_verified";
