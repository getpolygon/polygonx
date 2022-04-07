-- Write your migrate up statements here
CREATE TABLE "upvotes" (
    "post"        CHAR(27)        NOT NULL REFERENCES posts("id"),
    "user"        CHAR(27)        NOT NULL REFERENCES users("id"),
    "created_at"  TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    PRIMARY KEY ("post", "user")
);
---- create above / drop below ----
DROP TABLE "upvotes";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
