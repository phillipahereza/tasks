package main

import (
	"os"
	"fmt"
	"github.com/phillipahereza/tasks/db"
	"path/filepath"
	"github.com/mitchellh/go-homedir"
	"github.com/phillipahereza/tasks/cmd"
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
