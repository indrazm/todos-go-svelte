package database

import (
	"todolist/server/models"
)

var TodoList []models.Task = []models.Task{
	{Id: 1, Title: "Get Project 1", IsCompleted: false},
	{Id: 2, Title: "Get Project 2", IsCompleted: false},
	{Id: 3, Title: "Get Project 3", IsCompleted: true},
}