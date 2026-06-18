CREATE TABLE profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    bio TEXT NOT NULL,
    avatar_url TEXT,
    city VARCHAR(100),
    looking_for TEXT[],
    skills TEXT[],
    available_for_projects BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now()
);