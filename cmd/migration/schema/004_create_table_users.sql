-- +goose Up
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "full_name" varchar(64) NOT NULL,
    "phone_number" varchar(64) NOT NULL,
    "birth_date" date NOT NULL,
    "photo" varchar(64) DEFAULT '',
    "role_id" bigint NOT NULL,
    "username" varchar(64) NOT NULL,
    "password" text NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    FOREIGN KEY ("role_id") REFERENCES "roles"("id")
);

CREATE UNIQUE INDEX unique_username_not_deleted ON "users" ("username") WHERE "deleted_at" IS NULL;

INSERT INTO "users" (
    "full_name",
    "phone_number",
    "birth_date",
    "role_id",
    "username",
    "password"
) VALUES (
    'Super Admin',
    '+998901234567',
    CURRENT_DATE::DATE,
    (SELECT "id" FROM "roles" WHERE "name" = 'SUPER ADMIN'),
    'superadmin',
    '38393071776572747975696f704153444647484a8c3e0a3a7dd0467d06dcad504c73e2b780cee92d'
);

-- +goose Down
DROP TABLE "users";