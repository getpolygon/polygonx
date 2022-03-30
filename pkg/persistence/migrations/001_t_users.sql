-- Write your migrate up statements here
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR NOT NULL,
    "password" VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    username VARCHAR NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
