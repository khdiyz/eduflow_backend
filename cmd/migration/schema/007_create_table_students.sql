-- +goose Up
CREATE TABLE "students" (
    "id" bigserial PRIMARY KEY,
    "full_name" varchar(64) NOT NULL,
    "phone_number_1" varchar(64) NOT NULL,
    "phone_number_2" varchar(64),
    "address" text,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp
);

-- +goose Down
DROP TABLE "students";