-- Table Definition
CREATE TABLE expenses
(
    id     SERIAL PRIMARY KEY,
    title  TEXT,
    amount FLOAT,
    note   TEXT,
    tags   TEXT[]
);

-- Seeding Data
INSERT INTO expenses (title, amount, note, tags)
VALUES ('strawberry smoothie', 79, 'night market promotion discount 10 bath', '{"food","beverage"}');