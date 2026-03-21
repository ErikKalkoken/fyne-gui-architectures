package main

import (
	"database/sql"
	"flag"
	"log"

	"fyne.io/fyne/v2/app"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/model"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/view"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/viewmodel"
)

var errorModeFlag = flag.Bool("error-mode", false, "When activated adding task will generate an error.")

func main() {
	flag.Parse()
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	m := model.New(db, *errorModeFlag)
	err = m.Init()
	if err != nil {
		log.Fatal(err)
	}
	a := app.New()
	vm := viewmodel.New(m)
	v := view.New(a, vm)
	vm.Init()
	v.Run()
}
