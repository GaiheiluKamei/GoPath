CREATE TABLE IF NOT EXISTS links (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    url STRING UNIQUE,
    retrieved_at TIMESTAMP
);