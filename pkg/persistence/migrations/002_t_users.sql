-- Write your migrate up statements here
CREATE TABLE "users" (
    "id"            CHAR(27)        NOT NULL PRIMARY KEY DEFAULT generate_ulid(),
    "name"          VARCHAR         NOT NULL,
    -- an email address can contain at most 254 characters
	-- more info: https://stackoverflow.com/questions/386294/what-is-the-maximum-length-of-a-valid-email-address#:~:text=%22There%20is%20a%20length%20limit,total%20length%20of%20320%20characters.
    "email"         VARCHAR(254)    NOT NULL UNIQUE,
    "password"      VARCHAR         NOT NULL,
    "username"      VARCHAR         NOT NULL
);
---- create above / drop below ----
DROP TABLE "users";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
