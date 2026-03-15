package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	addButton    *widget.Button
	app          fyne.App
	deleteButton *widget.Button
	entry        *widget.Entry
	selected     string
	taskList     *widget.List
	tasks        []string
	w            fyne.Window
}

func NewView(a fyne.App) *View {
	v := &View{
		app:   a,
		tasks: make([]string, 0),
	}
	v.w = v.app.NewWindow("ToDo List")
	return v
}

func (v *View) MakeUI(presenter *Presenter) {
	v.addButton = widget.NewButton("Add", func() {
		go presenter.OnAddTask(v.entry.Text)
	})
	v.addButton.Disable()

	v.entry = widget.NewEntry()
	v.entry.PlaceHolder = "Name of new task"
	v.entry.OnChanged = func(s string) {
		if len(s) > 0 {
			v.addButton.Enable()
		} else {
			v.addButton.Disable()
		}
	}

	v.deleteButton = widget.NewButton("Delete", func() {
		go presenter.OnDeleteTask(v.selected)
	})
	v.deleteButton.Disable()

	v.taskList = widget.NewList(
		func() int {
			return len(v.tasks)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			x := v.tasks[id]
			co.(*widget.Label).SetText(x)
		},
	)
	v.taskList.OnSelected = func(id widget.ListItemID) {
		x := v.tasks[id]
		v.selected = x
		v.deleteButton.Enable()
	}
	v.taskList.OnUnselected = func(id widget.ListItemID) {
		v.selected = ""
		v.deleteButton.Disable()
	}

	c := container.NewBorder(
		nil,
		container.NewVBox(
			container.NewBorder(nil, nil, nil, v.addButton, v.entry),
			v.deleteButton,
		),
		nil,
		nil,
		v.taskList,
	)
	v.w.SetContent(c)
	v.w.Resize(fyne.NewSize(300, 500))
}

func (v *View) ShowError(err error) {
	dialog.ShowError(err, v.w)
}

func (v *View) ClearEntry() {
	v.entry.SetText("")
}

func (v *View) UpdateTaskList(tasks []string) {
	v.tasks = tasks
	v.taskList.Refresh()
}

func (v *View) Run() {
	v.w.ShowAndRun()
}
