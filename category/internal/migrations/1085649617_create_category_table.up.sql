CREATE TABLE IF NOT EXISTS category (
    id UUID NOT NULL,
    name varchar(100) NOT NULL,
    created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX name ON category(name)