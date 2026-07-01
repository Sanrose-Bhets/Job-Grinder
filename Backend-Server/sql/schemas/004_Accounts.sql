-- +goose up
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    Role VARCHAR(100) NOT NULL default 'user',
    APIKey VARCHAR(64) UNIQUE NOT NULL DEFAULT(
        encode(sha256(random()::text::bytea),'hex')
    )
    
);
-- +goose Down
DROP TABLE accounts;