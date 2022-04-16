-- Write your migrate up statements here
CREATE TABLE "comments" (
    "id" CHAR(27) NOT NULL DEFAULT generate_ulid(),
    "post" CHAR(27) NOT NULL REFERENCES "posts"("id"),
    "user" CHAR(27) NOT NULL REFERENCES "users"("id"),
    "content" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

---- create above / drop below ----
DROP TABLE "comments";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.