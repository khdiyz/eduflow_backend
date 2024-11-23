-- +goose Up
CREATE TABLE IF NOT EXISTS "schools" (
    "id" uuid PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "address" text,
    "email" varchar(255),
    "phone_number" varchar(64),
    "currency" varchar(8) NOT NULL,
    "timezone" varchar(64) NOT NULL DEFAULT 'Asia/Tashkent',
    "status" boolean NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp
);

-- +goose Down
DROP TABLE IF EXISTS "schools";
