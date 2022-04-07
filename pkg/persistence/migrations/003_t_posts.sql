-- Write your migrate up statements here
CREATE TABLE "posts" (
    "id"            CHAR(27)        NOT NULL PRIMARY KEY DEFAULT generate_ulid(),
    "user"          CHAR(27)        NOT NULL REFERENCES users("id"),
    "title"         VARCHAR(120)    NOT NULL,
    "content"       TEXT                NULL,
    "updated"       BOOLEAN         NOT NULL DEFAULT FALSE,
    "updated_at"    TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

---- create above / drop below ----
DROP TABLE "posts";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
