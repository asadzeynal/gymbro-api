CREATE TYPE "body_part" AS ENUM (
  'back',
  'chest',
  'other'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar,
  "fullName" varchar,
  "imageUrl" varchar
);

CREATE TABLE "dict_exercises" (
  "id" smallint,
  "body_part" body_part,
  "description" text,
  "imageUrl" varchar
);

CREATE TABLE "workouts_performed" (
  "id" uuid,
  "user_id" uuid,
  "name" varchar,
  "started_at" timestamp,
  "completed_at" timestamp
);

CREATE TABLE "exercises_performed" (
  "id" uuid,
  "exercise_id" smallint,
  "workout_performed_id" uuid
);

CREATE TABLE "sets_performed" (
  "exercise_performed_id" uuid,
  "weight" decimal,
  "reps" smallint
);

ALTER TABLE "workouts_performed" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "exercises_performed" ADD FOREIGN KEY ("exercise_id") REFERENCES "dict_exercises" ("id");

ALTER TABLE "exercises_performed" ADD FOREIGN KEY ("workout_performed_id") REFERENCES "workouts_performed" ("id");

ALTER TABLE "sets_performed" ADD FOREIGN KEY ("exercise_performed_id") REFERENCES "exercises_performed" ("id");
