-- +migrate Up
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    username  VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email CHAR(100) UNIQUE NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
