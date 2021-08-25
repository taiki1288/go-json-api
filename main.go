package main

import (
	"time"
) 

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"due_date"`
}

var tasks = []Task{{
	ID: 1,
	Title: "A",
	Content: "Aタスク",
	DueDate: time.Now(),
}, {
	ID: 2,
	Title: "B",
	Content: "Bタスク",
	DueDate: time.Now(),
}, {
	ID: 3,
	Title: "C",
	Content: "Cタスク",
	DueDate: time.Now(),
}}