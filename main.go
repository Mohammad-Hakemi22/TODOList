package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"
)

func main() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(homedir, "tasks.db")
	fmt.Println(dbPath)
	must(db.InitDb(dbPath))
	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
