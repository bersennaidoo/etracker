// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: exercise.sql

package pgstore

import (
	"context"
	"time"
)

const createDefaultSetForExercise = `-- name: CreateDefaultSetForExercise :one
INSERT INTO etrackerapp.sets (
    Workout_ID,
    Exercise_Name,
    Weight
) VALUES (
    $1,
    $2,
    $3
) RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type CreateDefaultSetForExerciseParams struct {
	WorkoutID    int64  `db:"workout_id" json:"workoutId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
	Weight       int32  `db:"weight" json:"weight"`
}

func (q *Queries) CreateDefaultSetForExercise(ctx context.Context, arg CreateDefaultSetForExerciseParams) (EtrackerappSet, error) {
	row := q.queryRow(ctx, q.createDefaultSetForExerciseStmt, createDefaultSetForExercise, arg.WorkoutID, arg.ExerciseName, arg.Weight)
	var i EtrackerappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}

const createSetForExercise = `-- name: CreateSetForExercise :one
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
) RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type CreateSetForExerciseParams struct {
	WorkoutID    int64  `db:"workout_id" json:"workoutId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
	Weight       int32  `db:"weight" json:"weight"`
	Set1         int64  `db:"set1" json:"set1"`
	Set2         int64  `db:"set2" json:"set2"`
	Set3         int64  `db:"set3" json:"set3"`
}

func (q *Queries) CreateSetForExercise(ctx context.Context, arg CreateSetForExerciseParams) (EtrackerappSet, error) {
	row := q.queryRow(ctx, q.createSetForExerciseStmt, createSetForExercise,
		arg.WorkoutID,
		arg.ExerciseName,
		arg.Weight,
		arg.Set1,
		arg.Set2,
		arg.Set3,
	)
	var i EtrackerappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}

const createUserDefaultExercise = `-- name: CreateUserDefaultExercise :exec
INSERT INTO etrackerapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    1,
    'Bench Press'
),(
    1,
    'Barbell Row'
)
`

func (q *Queries) CreateUserDefaultExercise(ctx context.Context) error {
	_, err := q.exec(ctx, q.createUserDefaultExerciseStmt, createUserDefaultExercise)
	return err
}

const createUserExercise = `-- name: CreateUserExercise :one
INSERT INTO etrackerapp.exercises (
    User_ID,
    Exercise_Name
) VALUES (
    $1,
    $2
) ON CONFLICT (Exercise_Name) DO NOTHING RETURNING (
    User_ID, Exercise_Name
)
`

type CreateUserExerciseParams struct {
	UserID       int64  `db:"user_id" json:"userId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
}

func (q *Queries) CreateUserExercise(ctx context.Context, arg CreateUserExerciseParams) (interface{}, error) {
	row := q.queryRow(ctx, q.createUserExerciseStmt, createUserExercise, arg.UserID, arg.ExerciseName)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const createUserWorkout = `-- name: CreateUserWorkout :one
INSERT INTO etrackerapp.workouts (
    User_ID,
    Start_Date
) VALUES (
    $1,
    NOW()
) RETURNING workout_id, user_id, start_date
`

func (q *Queries) CreateUserWorkout(ctx context.Context, userID int64) (EtrackerappWorkout, error) {
	row := q.queryRow(ctx, q.createUserWorkoutStmt, createUserWorkout, userID)
	var i EtrackerappWorkout
	err := row.Scan(&i.WorkoutID, &i.UserID, &i.StartDate)
	return i, err
}

const deleteUserExercise = `-- name: DeleteUserExercise :exec
DELETE FROM etrackerapp.exercises
WHERE User_ID = $1 AND Exercise_Name = $2
`

type DeleteUserExerciseParams struct {
	UserID       int64  `db:"user_id" json:"userId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
}

func (q *Queries) DeleteUserExercise(ctx context.Context, arg DeleteUserExerciseParams) error {
	_, err := q.exec(ctx, q.deleteUserExerciseStmt, deleteUserExercise, arg.UserID, arg.ExerciseName)
	return err
}

const deleteWorkoutByIDForUser = `-- name: DeleteWorkoutByIDForUser :exec
DELETE FROM etrackerapp.workouts
WHERE User_ID = $1 AND Workout_ID = $2
`

type DeleteWorkoutByIDForUserParams struct {
	UserID    int64 `db:"user_id" json:"userId"`
	WorkoutID int64 `db:"workout_id" json:"workoutId"`
}

func (q *Queries) DeleteWorkoutByIDForUser(ctx context.Context, arg DeleteWorkoutByIDForUserParams) error {
	_, err := q.exec(ctx, q.deleteWorkoutByIDForUserStmt, deleteWorkoutByIDForUser, arg.UserID, arg.WorkoutID)
	return err
}

const getWorkoutsForUserID = `-- name: GetWorkoutsForUserID :many
SELECT w.Workout_ID, COALESCE(s.Set_ID,-1), COALESCE(s.name,''), COALESCE(s.set1,-1), COALESCE(s.set1,-1), COALESCE(s.set2,-1), COALESCE(s.set3,-1), COALESCE(s.weight,-1), w.Start_Date AS date FROM
(
SELECT Set_ID, Workout_ID, Exercise_Name as name, set1, set2, set3, weight FROM etrackerapp.sets
) AS s RIGHT JOIN etrackerapp.workouts AS w USING (Workout_ID)
WHERE w.User_ID = $1
ORDER BY date DESC
`

type GetWorkoutsForUserIDRow struct {
	WorkoutID int64     `db:"workout_id" json:"workoutId"`
	SetID     int64     `db:"set_id" json:"setId"`
	Name      string    `db:"name" json:"name"`
	Set1      int64     `db:"set1" json:"set1"`
	Set1_2    int64     `db:"set1_2" json:"set12"`
	Set2      int64     `db:"set2" json:"set2"`
	Set3      int64     `db:"set3" json:"set3"`
	Weight    int32     `db:"weight" json:"weight"`
	Date      time.Time `db:"date" json:"date"`
}

func (q *Queries) GetWorkoutsForUserID(ctx context.Context, userID int64) ([]GetWorkoutsForUserIDRow, error) {
	rows, err := q.query(ctx, q.getWorkoutsForUserIDStmt, getWorkoutsForUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetWorkoutsForUserIDRow
	for rows.Next() {
		var i GetWorkoutsForUserIDRow
		if err := rows.Scan(
			&i.WorkoutID,
			&i.SetID,
			&i.Name,
			&i.Set1,
			&i.Set1_2,
			&i.Set2,
			&i.Set3,
			&i.Weight,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserExercises = `-- name: ListUserExercises :many
SELECT Exercise_Name
FROM etrackerapp.exercises
WHERE User_ID = $1
`

func (q *Queries) ListUserExercises(ctx context.Context, userID int64) ([]string, error) {
	rows, err := q.query(ctx, q.listUserExercisesStmt, listUserExercises, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var exercise_name string
		if err := rows.Scan(&exercise_name); err != nil {
			return nil, err
		}
		items = append(items, exercise_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSet = `-- name: UpdateSet :one
UPDATE etrackerapp.sets SET
    Weight = $1,
    Set1 = $2,
    Set2 = $3,
    Set3 = $4
WHERE Set_ID = $5 AND Workout_ID = $6 RETURNING set_id, workout_id, exercise_name, weight, set1, set2, set3
`

type UpdateSetParams struct {
	Weight    int32 `db:"weight" json:"weight"`
	Set1      int64 `db:"set1" json:"set1"`
	Set2      int64 `db:"set2" json:"set2"`
	Set3      int64 `db:"set3" json:"set3"`
	SetID     int64 `db:"set_id" json:"setId"`
	WorkoutID int64 `db:"workout_id" json:"workoutId"`
}

func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) (EtrackerappSet, error) {
	row := q.queryRow(ctx, q.updateSetStmt, updateSet,
		arg.Weight,
		arg.Set1,
		arg.Set2,
		arg.Set3,
		arg.SetID,
		arg.WorkoutID,
	)
	var i EtrackerappSet
	err := row.Scan(
		&i.SetID,
		&i.WorkoutID,
		&i.ExerciseName,
		&i.Weight,
		&i.Set1,
		&i.Set2,
		&i.Set3,
	)
	return i, err
}
