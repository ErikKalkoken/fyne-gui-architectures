package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvc/model"
)

type ToDoList struct {
	addButton    *widget.Button
	app          fyne.App
	deleteButton *widget.Button
	entry        *widget.Entry
	model        *model.Model
	onAddTask    func(string)
	onDeleteTask func(string)
	selected     string
	taskList     *widget.List
	tasks        []string
	w            fyne.Window
}

func NewToDoList(a fyne.App, model *model.Model) *ToDoList {
	v := &ToDoList{
		app:   a,
		tasks: make([]string, 0),
		model: model,
	}
	v.w = v.app.NewWindow("ToDo List")
	v.createUI()
	return v
}

func (v *ToDoList) createUI() {
	v.addButton = widget.NewButton("Add", func() {
		v.onAddTask(v.entry.Text)
	})
	v.addButton.Disable()

	v.entry = widget.NewEntry()
	v.entry.OnChanged = func(s string) {
		if len(s) > 0 {
			v.addButton.Enable()
		} else {
			v.addButton.Disable()
		}
	}

	v.deleteButton = widget.NewButton("Delete", func() {
		if x := v.selected; x != "" {
			v.onDeleteTask(x)
		}
		v.taskList.UnselectAll()
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

func (v *ToDoList) Init() error {
	err := v.UpdateTaskList()
	if err != nil {
		return err
	}
	return nil
}

func (v *ToDoList) BindAddTask(f func(string)) {
	v.onAddTask = f
}

func (v *ToDoList) BindDeleteTask(f func(string)) {
	v.onDeleteTask = f
}

func (v *ToDoList) ClearEntry() {
	v.entry.SetText("")
}

func (v *ToDoList) UpdateTaskList() error {
	tasks, err := v.model.ListTasks()
	if err != nil {
		return err
	}
	fyne.Do(func() {
		v.tasks = tasks
		v.taskList.Refresh()
	})
	return nil
}

func (v *ToDoList) Run() {
	v.w.ShowAndRun()
}
