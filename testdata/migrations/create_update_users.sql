CREATE TABLE "update_user" (
    id SERIAL PRIMARY KEY,
    user_id INT,
    name NVARCHAR(255),
    email NVARCHAR(255),
    updated_at DATE
    FOREIGN KEY (user_id) REFERENCES "users" (id)
);
