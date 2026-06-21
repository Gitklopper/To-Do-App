package main

import "time"

type Progression string

const (
	ProgressionOpen			Progression = "open"
	ProgressionInProgress	Progression = "in_progress"
	ProgressionDone			Progression = "done"
)

type Priority string

const (
	PriorityUrgantImportant		Priority = "urgent_important"
	PriorityUrgent				Priority = "urgent"
	PriorityImportant			Priority = "important"
	PriorityNeither 			Priority = "neither"
)

type ToDo struct {
	ID 			int			`json:"id"`
	Title 		string		`json:"title"`
	Status 		Progression	`json:"progression"`
	Role 		Priority	`json:"priority"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}
