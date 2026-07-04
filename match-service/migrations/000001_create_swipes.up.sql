CREATE TABLE swipes (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_user   UUID NOT NULL,
    to_user     UUID NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    UNIQUE (from_user, to_user)
);

CREATE INDEX idx_swipes_from_user ON swipes(from_user);
CREATE INDEX idx_swipes_to_user ON swipes(to_user);