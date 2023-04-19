CREATE TABLE transactions
(
    id         SERIAL PRIMARY KEY,
    user_id    int            NOT NULL,
    amount     DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    book_name  VARCHAR(255)
);