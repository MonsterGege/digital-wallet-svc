CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE wallets (
    wallet_id INT PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id),
    balance DECIMAL(15,2) NOT NULL DEFAULT 0
);

CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    wallet_id INT REFERENCES wallets(wallet_id),
    amount DECIMAL(15,2) NOT NULL,
    types VARCHAR(20) NOT NULL, -- withdraw / deposit
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (name) VALUES ('Alice'), ('Bob');
INSERT INTO wallets (wallet_id, user_id, balance) VALUES (1001, 1, 1000.00), (1002, 2, 500.00);
-- Sample transactions
INSERT INTO transactions (wallet_id, amount, types) VALUES
(1001, 200.00, 'withdraw'),
(1001, 300.00, 'deposit'),
(1002, 100.00, 'withdraw');