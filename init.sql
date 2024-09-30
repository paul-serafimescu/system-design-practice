-- Create users
DO $$
BEGIN
    -- Create wss_dev if not exists
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'wss_dev') THEN
        CREATE ROLE wss_dev WITH LOGIN PASSWORD 'postgres';
    END IF;

    -- Create user2 if not exists
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'api_dev') THEN
        CREATE ROLE api_dev WITH LOGIN PASSWORD 'postgres';
    END IF;
END $$;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/************************************/
/* START HTTP SERVER                */
/************************************/

CREATE TABLE IF NOT EXISTS Users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash CHAR(64) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    firstname VARCHAR(50),
    lastname VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON Users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();


-- DUMMY DATA --
INSERT INTO Users (user_id, username, password_hash, email, firstname, lastname)
VALUES ('644dae0b-0287-490d-abc2-79fc64d7ae0f', 'sample-username', '8dccd2fa9e2340d455740d1b4ae2dca6774c3b7d7dbdef9f362c1c59ecec4016', 'noreply@gmail.com', 'Paul', 'Serafimescu');

/************************************/
/* END HTTP SERVER                  */
/************************************/

-- Grant privileges to users
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO wss_dev;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO api_dev;
