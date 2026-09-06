CREATE TABLE books (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    author text NOT NULL,
    title text NOT NULL,
    copies int NOT NULL,
    price_cents int NOT NULL
);
