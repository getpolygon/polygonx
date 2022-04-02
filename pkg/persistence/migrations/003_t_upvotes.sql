-- Write your migrate up statements here
CREATE TABLE "upvotes" (
    "post"        UUID            NOT NULL REFERENCES posts("id"),
    "user"        UUID            NOT NULL REFERENCES users("id"),
    "created_at"  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("post", "user")
);

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
