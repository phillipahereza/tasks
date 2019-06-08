package db

import (
	"github.com/asdine/storm"
	// bolt "go.etcd.io/bbolt"
	"log"
	// "time"
)

var db *storm.DB

// Task rep
type Task struct {
	ID        int `storm:"id,increment"`
	Value     string
}

// Init initializes database
func Init(dbPath string) error {
	var err error
	db, err = storm.Open(dbPath)
	return err
}

// CreateTask adds task to the DB
func CreateTask(value string) (Task, error) {
	task := Task{Value: value}
	err := db.Save(&task)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return task, err
}

// FetchTasks retrieves all tasks
func FetchTasks() ([]Task, error){
	var tasks []Task
	err := db.All(&tasks)
	
	return tasks, err
}

// DoTask marks a task as complete
func DoTask(id int) error {
	err := db.Delete("Task", id)
	if err != nil {
		return err
	}

	return db.ReIndex(&Task{})
}

// TODO add deleted, timeCreated, timeCompleted to Task
// TODO add remove/delete a task
// TODO add complete command to show Tasks completed within 12hrs 24hrs, i.e use flags
// TODO add ability to track time of task, start stop track
// TODO show status of tasks

