CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(100),
    bio TEXT,
    tags TEXT[],
    looking_for TEXT[],
    last_seen TIMESTAMP DEFAULT NOW()
);