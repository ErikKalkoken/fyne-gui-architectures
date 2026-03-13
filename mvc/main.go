package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvc/tasks"
)

func main() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	m := tasks.NewModel(db)
	err = m.Init()
	if err != nil {
		log.Fatal(err)
	}
	a := app.New()
	v := tasks.NewView(a, m)
	err = v.Init()
	if err != nil {
		log.Fatal(err)
	}
	c := tasks.NewController(m, v)
	c.Run()
}
