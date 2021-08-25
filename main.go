package main

import (
	"time"
) 

type Task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	DueDate time.Time `json:"due_date"`
}