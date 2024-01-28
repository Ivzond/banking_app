CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name TEXT,
    username TEXT,
    email TEXT,
    password TEXT
);

ALTER TABLE users OWNER TO postgres;

CREATE TABLE IF NOT EXISTS accounts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    type TEXT,
    name TEXT,
    balance INTEGER,
    user_id INTEGER REFERENCES users
);

ALTER TABLE accounts OWNER TO postgres;

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    "from" INTEGER REFERENCES accounts,
    "to" INTEGER REFERENCES accounts,
    amount INTEGER
);

ALTER TABLE transactions OWNER TO postgres;
