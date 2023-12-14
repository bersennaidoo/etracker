-- name: CreateUserExercise :one
INSERT INTO etrackerapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    $1,
    $2
) ON CONFLICT (Exercise_Name) DO NOTHING RETURNING (
    User_ID, Exercise_Name
);

-- name: ListUserExercises :many
SELECT Exercise_Name
FROM etrackerapp.exercises
WHERE User_ID = $1;

-- name: DeleteUserExercise :exec
DELETE FROM etrackerapp.exercises
WHERE User_ID = $1 AND Exercise_Name = $2;


-- name: CreateUserDefaultExercise :exec
INSERT INTO etrackerapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    1,
    'Bench Press'
),(
    1,
    'Barbell Row'
);

-- name: CreateUserWorkout :one
INSERT INTO etrackerapp.workouts (
    User_ID,
    Start_Date
) VALUES (
    $1,
    NOW()
) RETURNING *;

-- name: CreateDefaultSetForExercise :one
INSERT INTO etrackerapp.sets (
    Workout_ID,
    Exercise_Name,
    Weight
) VALUES (
    $1,
    $2,
    $3
) RETURNING *;

-- name: CreateSetForExercise :one
INSERT INTO etrackerapp.sets (
    Workout_ID,
    Exercise_Name, 
    Weight,
    Set1,
    Set2,
    Set3
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: UpdateSet :one
UPDATE etrackerapp.sets SET
    Weight = $1,
    Set1 = $2,
    Set2 = $3,
    Set3 = $4
WHERE Set_ID = $5 AND Workout_ID = $6 RETURNING *;

-- name: GetWorkoutsForUserID :many
SELECT w.Workout_ID, COALESCE(s.Set_ID,-1), COALESCE(s.name,''), COALESCE(s.set1,-1), COALESCE(s.set1,-1), COALESCE(s.set2,-1), COALESCE(s.set3,-1), COALESCE(s.weight,-1), w.Start_Date AS date FROM
(
SELECT Set_ID, Workout_ID, Exercise_Name as name, set1, set2, set3, weight FROM etrackerapp.sets
) AS s RIGHT JOIN etrackerapp.workouts AS w USING (Workout_ID)
WHERE w.User_ID = $1
ORDER BY date DESC;


-- name: DeleteWorkoutByIDForUser :exec
DELETE FROM etrackerapp.workouts
WHERE User_ID = $1 AND Workout_ID = $2;
