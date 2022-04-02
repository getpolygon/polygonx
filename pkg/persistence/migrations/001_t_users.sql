-- Write your migrate up statements here
CREATE TABLE "users" (
    "id"            UUID            NOT NULL PRIMARY KEY,
    "name"          VARCHAR         NOT NULL,
    "email"         VARCHAR         NOT NULL UNIQUE,
    "password"      VARCHAR         NOT NULL,
    "username"      VARCHAR         NOT NULL,
    "created_at"    TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
