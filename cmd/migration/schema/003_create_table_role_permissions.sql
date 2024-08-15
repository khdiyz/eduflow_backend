-- +goose Up
CREATE TABLE "role_permissions" (
    "id" bigserial PRIMARY KEY,
    "role_id" bigint NOT NULL,
    "permission_id" bigint NOT NULL,
    "status" boolean NOT NULL,
    FOREIGN KEY ("role_id") REFERENCES "roles"("id"),
    FOREIGN KEY ("permission_id") REFERENCES "permissions"("id")
);

-- +goose Down
DROP TABLE "role_permissions";