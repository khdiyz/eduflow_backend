-- +goose Up
CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY,
    "role_id" uuid NOT NULL,
    "first_name" varchar(64) NOT NULL,
    "last_name" varchar(64),
    "phone_numbers" varchar(64)[],
    "username" varchar(64) NOT NULL UNIQUE,
    "password" text NOT NULL,
    "status" boolean NOT NULL,
    "branch_id" uuid,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    FOREIGN KEY (branch_id) REFERENCES branches(id)
);

-- +goose Down
DROP TABLE IF EXISTS "users";
