-- Write your migrate up statements here
CREATE TABLE comments (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    post_id UUID NOT NULL REFERENCES posts(id),
    user_id UUID NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
