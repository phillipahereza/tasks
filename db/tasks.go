package db

import (
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
	// bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

var db *storm.DB

// Task rep
type Task struct {
	ID        int `storm:"id,increment"`
	Value     string
	Completed bool
	Deleted   bool
	CreatedAt time.Time
}

// Init initializes database
func Init(dbPath string) error {
	var err error
	db, err = storm.Open(dbPath)
	return err
}

// CreateTask adds task to the DB
func CreateTask(value string) (Task, error) {
	task := Task{Value: value, Completed: false, Deleted: false, CreatedAt: time.Now()}
	err := db.Save(&task)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return task, err
}

// FetchTasks retrieves all tasks
func FetchTasks() ([]Task, error) {
	var tasks []Task
	err := db.Select(q.And(q.Eq("Completed", false), q.Eq("Deleted", false))).Find(&tasks)
	return tasks, err
}

// DoTask marks a task as complete
func DoTask(id int) error {
	return db.UpdateField(&Task{ID: id}, "Completed", true)
}

// DeleteTask marks a task as deleted
func DeleteTask(id int) error {
	return db.UpdateField(&Task{ID: id}, "Deleted", true)
}

// TODO add deleted, timeCreated, timeCompleted to Task
// TODO add remove/delete a task
// TODO add complete command to show Tasks completed within 12hrs 24hrs, i.e use flags
// TODO add ability to track time of task, start stop track
// TODO show status of tasks
