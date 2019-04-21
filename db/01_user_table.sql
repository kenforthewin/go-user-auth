-- rambler up

CREATE TABLE users (
id BIGSERIAL PRIMARY KEY,
email VARCHAR(255) UNIQUE NOT NULL,
created_at TIMESTAMP NOT NULL,
password VARCHAR(255) NOT NULL
);

-- rambler down

DROP TABLE users;
