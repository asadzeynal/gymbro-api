Table exercises {
  id uuid [pk]
  name varchar [not null]
  description varchar
}

Table users {
  id uuid [primary key]
  email varchar(255) [not null, unique]
  password varchar(255) [not null]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table templates {
  id uuid [primary key]
  name varchar(255) [not null]
  user_id uuid [not null, ref: > users.id]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table template_exercises {
  id uuid [primary key]
  template_id uuid [not null, ref: > templates.id]
  exercise_id varchar(255) [not null, ref: > exercises.id]
  display_order int [not null, default: 0]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table template_sets {
  id uuid [primary key]
  template_exercise_id uuid [not null, ref: > template_exercises.id]
  reps int [not null, default: 0]
  weight float [not null, default: 0.0]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table workouts {
  id uuid [primary key]
  user_id uuid [not null, ref: > users.id]
  name varchar(255) [not null]
  date timestamp [not null, default: `now()`]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table workout_exercises {
  id uuid [primary key]
  workout_id uuid [not null, ref: > workouts.id]
  exercise_id varchar(255) [not null, ref: > exercises.id]
  display_order int [not null, default: 0]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}

Table workout_sets {
  id uuid [primary key]
  workout_exercise_id uuid [not null, ref: > workout_exercises.id]
  reps int [not null, default: 0]
  weight float [not null, default: 0.0]
  created_at timestamp [not null, default: `now()`]
  updated_at timestamp [not null, default: `now()`]
}
