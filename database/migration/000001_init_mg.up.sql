CREATE TABLE "users" (
                         "id" bigserial PRIMARY KEY,
                         "phone_number" varchar(15) NOT NULL,
                         "email" varchar(255) NOT NULL UNIQUE,
                         "username" varchar(50) NOT NULL UNIQUE,
                         "password" varchar(255) NOT NULL,
                         "role" varchar(20) NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

-- Ensure the email and username are unique
CREATE UNIQUE INDEX ON "users" ("email");
CREATE UNIQUE INDEX ON "users" ("username");

-- Add check constraint to ensure role is one of admin, seller, or buyer
ALTER TABLE "users"
    ADD CONSTRAINT "check_role"
        CHECK ("role" IN ('admin', 'seller', 'buyer'));

-- Indexes for common lookup fields
CREATE INDEX ON "users" ("phone_number");
CREATE INDEX ON "users" ("role");

-- Comments for clarity
COMMENT ON COLUMN "users"."password" IS 'hashed password';
COMMENT ON COLUMN "users"."role" IS 'role-based access: admin, seller, or buyer';

-- -- Example trigger function to hash password before insert or update
-- CREATE OR REPLACE FUNCTION hash_password_function()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   IF TG_OP = 'INSERT' OR (TG_OP = 'UPDATE' AND NEW.password IS DISTINCT FROM OLD.password) THEN
--     NEW.password = crypt(NEW.password, gen_salt('bf'));
-- END IF;
-- RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;
--
-- -- Trigger to call the hash_password_function before inserting a new user
-- CREATE TRIGGER hash_password_before_insert
--     BEFORE INSERT ON "users"
--     FOR EACH ROW
--     EXECUTE FUNCTION hash_password_function();
--
-- -- Trigger to call the hash_password_function before updating a user
-- CREATE TRIGGER hash_password_before_update
--     BEFORE UPDATE ON "users"
--     FOR EACH ROW
--     EXECUTE FUNCTION hash_password_function();
