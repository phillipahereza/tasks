package db

import (
	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	// bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

var db *storm.DB

// Task rep
type Task struct {
	ID          int `storm:"id,increment"`
	Value       string
	Completed   bool
	Deleted     bool
	CreatedAt   time.Time
	CompletedAt time.Time
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
	return db.Update(&Task{ID: id, CompletedAt: time.Now(), Completed: true})
	// return db.UpdateField(&Task{ID: id}, "Completed", true)
}

// DeleteTask marks a task as deleted
func DeleteTask(id int) error {
	return db.UpdateField(&Task{ID: id}, "Deleted", true)
}

// FetchCompletedTasks returns a list of all completed tasks
func FetchCompletedTasks(timeDelta time.Duration) ([]Task, error) {
	var tasks []Task
	timeLimit := time.Now().Add(-timeDelta)
	err := db.Select(q.And(q.Eq("Completed", true), q.Eq("Deleted", false), q.Gt("CompletedAt", timeLimit))).Find(&tasks)
	return tasks, err
}

// TODO add ability to track time of task, start stop track
// TODO show status of tasks
// TODO format data into tables
