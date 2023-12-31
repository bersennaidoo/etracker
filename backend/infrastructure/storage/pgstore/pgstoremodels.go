// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package pgstore

import (
	"encoding/json"
	"time"
)

type EtrackerappExercise struct {
	UserID       int64  `db:"user_id" json:"userId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
}

type EtrackerappImage struct {
	ImageID     int64  `db:"image_id" json:"imageId"`
	UserID      int64  `db:"user_id" json:"userId"`
	ContentType string `db:"content_type" json:"contentType"`
	ImageData   []byte `db:"image_data" json:"imageData"`
}

type EtrackerappSet struct {
	SetID        int64  `db:"set_id" json:"setId"`
	WorkoutID    int64  `db:"workout_id" json:"workoutId"`
	ExerciseName string `db:"exercise_name" json:"exerciseName"`
	Weight       int32  `db:"weight" json:"weight"`
	Set1         int64  `db:"set1" json:"set1"`
	Set2         int64  `db:"set2" json:"set2"`
	Set3         int64  `db:"set3" json:"set3"`
}

type EtrackerappUser struct {
	UserID       int64           `db:"user_id" json:"userId"`
	UserName     string          `db:"user_name" json:"userName"`
	PasswordHash string          `db:"password_hash" json:"passwordHash"`
	Name         string          `db:"name" json:"name"`
	Config       json.RawMessage `db:"config" json:"config"`
	CreatedAt    time.Time       `db:"created_at" json:"createdAt"`
	IsEnabled    bool            `db:"is_enabled" json:"isEnabled"`
}

type EtrackerappWorkout struct {
	WorkoutID int64     `db:"workout_id" json:"workoutId"`
	UserID    int64     `db:"user_id" json:"userId"`
	StartDate time.Time `db:"start_date" json:"startDate"`
}
