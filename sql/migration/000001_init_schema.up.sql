CREATE TABLE "exercises" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar NOT NULL
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar(255) UNIQUE NOT NULL,
  "password" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "templates" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "template_exercises" (
  "id" uuid PRIMARY KEY,
  "template_id" uuid NOT NULL,
  "exercise_id" uuid NOT NULL,
  "display_order" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "template_sets" (
  "id" uuid PRIMARY KEY,
  "template_exercise_id" uuid NOT NULL,
  "reps" int NOT NULL DEFAULT 0,
  "weight" float NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "workouts" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "name" varchar(255) NOT NULL,
  "date" timestamp NOT NULL DEFAULT (now()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "workout_exercises" (
  "id" uuid PRIMARY KEY,
  "workout_id" uuid NOT NULL,
  "exercise_id" uuid NOT NULL,
  "display_order" int NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "workout_sets" (
  "id" uuid PRIMARY KEY,
  "workout_exercise_id" uuid NOT NULL,
  "reps" int NOT NULL DEFAULT 0,
  "weight" float NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "templates" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "template_exercises" ADD FOREIGN KEY ("template_id") REFERENCES "templates" ("id");

ALTER TABLE "template_exercises" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");

ALTER TABLE "template_sets" ADD FOREIGN KEY ("template_exercise_id") REFERENCES "template_exercises" ("id");

ALTER TABLE "workouts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("workout_id") REFERENCES "workouts" ("id");

ALTER TABLE "workout_exercises" ADD FOREIGN KEY ("exercise_id") REFERENCES "exercises" ("id");

ALTER TABLE "workout_sets" ADD FOREIGN KEY ("workout_exercise_id") REFERENCES "workout_exercises" ("id");
