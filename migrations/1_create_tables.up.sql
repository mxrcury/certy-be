-- +migrate Up
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    username  VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email CHAR(100) UNIQUE NOT NULL,

    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),

    language VARCHAR(10) NOT NULL DEFAULT 'en',
    photo TEXT,
    job_title VARCHAR(100),

    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
