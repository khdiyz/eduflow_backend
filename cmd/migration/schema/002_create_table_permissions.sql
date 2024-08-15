-- +goose Up
CREATE TABLE "permissions" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(64) NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp
);

-- +goose Down
DROP TABLE "permissions";