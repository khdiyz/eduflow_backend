-- +goose Up
CREATE TABLE "roles" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(64) NOT NULL,
    "description" text,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp
);

CREATE UNIQUE INDEX unique_role_name_not_deleted ON "roles" ("name") WHERE "deleted_at" IS NULL;

INSERT INTO "roles" ("name") VALUES 
('SUPER ADMIN'),
('ADMIN'),
('MENTOR');

-- +goose Down
DROP TABLE "roles";