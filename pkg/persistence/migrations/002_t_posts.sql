-- Write your migrate up statements here
CREATE TABLE "posts" (
    "id"            UUID            NOT NULL PRIMARY KEY,
    "user"          UUID            NOT NULL REFERENCES users("id"),
    "title"         VARCHAR(120)    NOT NULL,
    "content"       TEXT                NULL,
    "updated_at"    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    "created_at"    TIMESTAMPTZ     NOT NULL DEFAULT NOW()
)

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
