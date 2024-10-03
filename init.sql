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

CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash CHAR(64) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    firstname VARCHAR(50),
    lastname VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE RegisteredServiceStatus AS ENUM ('UP', 'DOWN');

CREATE TABLE IF NOT EXISTS registered_service (
    registered_service_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    hostname VARCHAR(255) NOT NULL,
    port INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    status RegisteredServiceStatus NOT NULL,
    type INT NOT NULL
);

ALTER TABLE registered_service
ALTER COLUMN Status TYPE RegisteredServiceStatus USING Status::RegisteredServiceStatus;

ALTER TABLE registered_service
ADD CONSTRAINT service_type_check CHECK (Type IN (0, 1));

CREATE TABLE IF NOT EXISTS community (
    community_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_name VARCHAR(32) NOT NULL,
    community_description VARCHAR(256),
    owner_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    region VARCHAR(32),
    icon_url VARCHAR(256),
    FOREIGN KEY (owner_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS channel (
    channel_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    channel_name VARCHAR(32) NOT NULL,
    channel_description VARCHAR(256),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    position INTEGER DEFAULT 0,
    community_id UUID NOT NULL,
    last_message_id UUID,
    FOREIGN KEY (community_id) REFERENCES community(community_id)
);

CREATE TABLE IF NOT EXISTS text_message (
    message_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    community_id UUID NOT NULL,
    channel_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content VARCHAR(2048) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    is_edit BOOLEAN DEFAULT FALSE,
    is_deleted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (community_id) REFERENCES community(community_id) ON DELETE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_services_updated_at
BEFORE UPDATE ON registered_service
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_community_updated_at
BEFORE UPDATE ON community
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_channel_updated_at
BEFORE UPDATE ON channel
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_text_message_updated_at
BEFORE UPDATE ON text_message
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
