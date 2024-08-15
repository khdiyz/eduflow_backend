-- +goose Up
CREATE TABLE "groups" (
    "id" bigserial PRIMARY KEY,
    "name" varchar(64) NOT NULL,
    "start_date" date NOT NULL DEFAULT CURRENT_DATE,
    "end_date" date,
    "course_id" bigint NOT NULL,
    "teacher_id" bigint NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    FOREIGN KEY ("course_id") REFERENCES "courses"("id"),
    FOREIGN KEY ("teacher_id") REFERENCES "users"("id")
);

-- +goose Down
DROP TABLE "groups";