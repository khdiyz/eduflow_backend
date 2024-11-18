-- +goose Up
CREATE TABLE IF NOT EXISTS "roles" (
    "id" UUID PRIMARY KEY,
    "name" JSONB NOT NULL,
    "description" JSONB NOT NULL
);

INSERT INTO "roles" ("id", "name", "description")
VALUES
(
    'f6c762a0-72ca-45f9-92f6-02f407a23d69',
    '{"en": "Super Admin", "ru": "Супер Админ", "uz": "Super Admin"}',
    '{
        "en": "The Super Admin has unrestricted access to all parts of the system and full control over all users.", 
        "ru": "Суперадминистратор имеет неограниченный доступ ко всем частям системы и полный контроль над всеми пользователями.", 
        "uz": "Super Admin tizimning barcha bo‘limlariga cheklanmagan kirish huquqiga ega va barcha foydalanuvchilar ustidan to‘liq nazoratni amalga oshiradi."
    }'
);

-- +goose Down
DROP TABLE IF EXISTS "roles";
