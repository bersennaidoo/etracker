-- name: ListUsers :many
-- get all users ordered by the username
SELECT *
FROM etrackerapp.users
ORDER BY user_name;

-- name: ListImages :many
-- get all images ordered by the id
SELECT *
FROM etrackerapp.images
ORDER BY image_id;

-- name: ListExercises :many
-- get all exercises ordered by the exercise name
SELECT *
FROM etrackerapp.exercises
ORDER BY exercise_name;

-- name: ListSets :many
-- get all exercise sets ordered by weight
SELECT *
FROM etrackerapp.sets
ORDER BY weight;

-- name: ListWorkouts :many
-- get all workouts ordered by id
SELECT *
FROM etrackerapp.workouts
ORDER BY workout_id;

-- name: GetUser :one
-- get users of a particular user_id
SELECT *
FROM etrackerapp.users
WHERE user_id = $1;

-- name: GetUserWorkout :many
-- get a particular user information and workouts
SELECT u.user_id, w.workout_id, w.start_date, w.set_id
FROM etrackerapp.users u,
     etrackerapp.workouts w
WHERE u.user_id = w.user_id
  AND u.user_id = $1;

-- name: GetUserSets :many
-- get a particular user information, exercise sets and workouts
SELECT u.user_id, w.workout_id, w.start_date, s.set_id, s.weight
FROM etrackerapp.users u,
     etrackerapp.workouts w,
     etrackerapp.sets s
WHERE u.user_id = w.user_id
  AND w.set_id = s.set_id
  AND u.user_id = $1;

-- name: GetUserImage :one
-- get a particular user image
SELECT u.name, u.user_id, i.image_data
FROM etrackerapp.users u,
     etrackerapp.images i
WHERE u.user_id = i.user_id
  AND u.user_id = $1;

-- name: DeleteUsers :exec
-- delete a particular user
DELETE
FROM etrackerapp.users
WHERE user_id = $1;

-- name: DeleteUserImage :exec
-- delete a particular user's image
DELETE
FROM etrackerapp.images i
WHERE i.user_id = $1;

-- name: DeleteUserWorkouts :exec
-- delete a particular user's workouts
DELETE
FROM etrackerapp.workouts w
WHERE w.user_id = $1;

-- name: DeleteExercise :exec
-- delete a particular exercise
DELETE
FROM etrackerapp.exercises e
WHERE e.exercise_id = $1;

-- name: DeleteSets :exec
-- delete a particular exercise sets
DELETE
FROM etrackerapp.sets s
WHERE s.set_id = $1;

-- name: CreateExercise :one
-- insert a new exercise
INSERT INTO etrackerapp.exercises (Exercise_Name)
values ($1) RETURNING Exercise_ID;

-- name: UpsertExercise :one
-- insert or update exercise of a particular id
INSERT INTO etrackerapp.exercises (Exercise_Name)
VALUES ($1) ON CONFLICT (Exercise_ID) DO
UPDATE
    SET Exercise_Name = EXCLUDED.Exercise_Name
    RETURNING Exercise_ID;

-- name: CreateUserImage :one
-- insert a new image
INSERT INTO etrackerapp.images (User_ID, Content_Type, Image_Data)
values ($1,
        $2,
        $3) RETURNING *;

-- name: UpsertUserImage :one
-- insert or update image of a particular id
INSERT INTO etrackerapp.images (Image_Data)
VALUES ($1) ON CONFLICT (Image_ID) DO
UPDATE
    SET Image_Data = EXCLUDED.Image_Data
    RETURNING Image_ID;


-- name: CreateSet :one
-- insert new exercise sets
INSERT INTO etrackerapp.sets (Exercise_Id, Weight)
values ($1,
        $2) RETURNING *;

-- name: UpdateSet :one
-- insert a sets id
UPDATE etrackerapp.sets
SET (Exercise_Id, Weight) = ($1, $2)
WHERE set_id = $3 RETURNING *;

-- name: CreateWorkout :one
-- insert new workouts
INSERT INTO etrackerapp.workouts (User_ID, Set_ID, Start_Date)
values ($1,
        $2,
        $3) RETURNING *;

-- name: UpsertWorkout :one
-- insert or update workouts based of a particular ID
INSERT INTO etrackerapp.workouts (User_ID, Set_ID, Start_Date)
values ($1,
        $2,
        $3) ON CONFLICT (Workout_ID) DO
UPDATE
    SET User_ID = EXCLUDED.User_ID,
    Set_ID = EXCLUDED.Set_ID,
    Start_Date = EXCLUDED.Start_Date
    RETURNING Workout_ID;

-- name: CreateUsers :one
-- insert new user
INSERT INTO etrackerapp.users (User_Name, Pass_Word_Hash, name)
VALUES ($1,
        $2,
        $3) RETURNING *;
