-- +goose Up
CREATE TABLE IF NOT EXISTS "branches" (
    "id" uuid PRIMARY KEY,
    "school_id" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "address" text,
    "email" varchar(255),
    "phone_number" varchar(64),
    "opening_hours" varchar(255) NOT NULL,
    "status" boolean NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    FOREIGN KEY (school_id) REFERENCES schools(id)
);

-- +goose Down
DROP TABLE IF EXISTS "branches";
