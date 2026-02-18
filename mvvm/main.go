package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/model"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/view"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/viewmodel"
)

func main() {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m := model.NewModel(db)
	err = m.Init()
	if err != nil {
		log.Fatal(err)
	}
	a := app.New()
	v := view.NewToDoList(a)
	p := viewmodel.NewViewModel(m, v)
	p.Run()
}
