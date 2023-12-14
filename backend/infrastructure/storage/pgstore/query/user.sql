-- name: ListUsers :many
SELECT *
FROM etrackerapp.users
ORDER BY user_name;

-- name: ListImages :many
SELECT *
FROM etrackerapp.images
ORDER BY image_id;


-- name: GetUser :one
SELECT *
FROM etrackerapp.users
WHERE user_id = $1;

-- name: GetUserByName :one
SELECT *
FROM etrackerapp.users
WHERE user_name = $1;

-- name: GetUserImage :one
SELECT u.name, u.user_id, i.image_data
FROM etrackerapp.users u,
     etrackerapp.images i
WHERE u.user_id = i.user_id
  AND u.user_id = $1;

-- name: DeleteUsers :exec
DELETE
FROM etrackerapp.users
WHERE user_id = $1;

-- name: DeleteUserImage :exec
DELETE
FROM etrackerapp.images i
WHERE i.user_id = $1;

-- name: DeleteUserWorkouts :exec
DELETE
FROM etrackerapp.workouts w
WHERE w.user_id = $1;


-- name: CreateUserImage :one
INSERT INTO etrackerapp.images (User_ID, Content_Type, Image_Data)
values ($1,
        $2,
        $3) RETURNING *;

-- name: UpsertUserImage :one
INSERT INTO etrackerapp.images (Image_Data)
VALUES ($1) ON CONFLICT (Image_ID) DO
UPDATE
    SET Image_Data = EXCLUDED.Image_Data
    RETURNING Image_ID;

-- name: CreateUsers :one
INSERT INTO etrackerapp.users (User_Name, Password_Hash, name)
VALUES ($1,
        $2,
        $3) RETURNING *;
