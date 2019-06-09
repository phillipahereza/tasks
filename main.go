package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/phillipahereza/tasks/cmd"
	"github.com/phillipahereza/tasks/db"
	"os"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, ".tasks.db")
	must(db.Init(dbPath))
	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
