package main

import (
	"database/sql"
	"flag"
	"log"

	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"
)

var errorModeFlag = flag.Bool("error-mode", false, "When activated adding task will generate an error.")

func main() {
	flag.Parse()
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m := NewModel(db, *errorModeFlag)
	err = m.Init()
	if err != nil {
		log.Fatal(err)
	}
	a := app.New()
	vm := NewViewModel(m)
	v := NewView(a, vm)
	v.Run()
}
