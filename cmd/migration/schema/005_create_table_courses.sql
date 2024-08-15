-- +goose Up
CREATE TABLE "courses" (
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(64) NOT NULL,
    "description" TEXT,
    "photo" VARCHAR(64),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP
);

-- +goose Down
DROP TABLE "courses";