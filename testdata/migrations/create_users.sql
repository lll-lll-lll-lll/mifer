CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    name NVARCHAR(255),
    email NVARCHAR(255),
    created_at DATE
);
