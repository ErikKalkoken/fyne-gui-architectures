package tasks

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
	model        *Model
	onAddTask    func(string)
	onDeleteTask func(string)
	selected     string
	taskList     *widget.List
	tasks        []string
	w            fyne.Window
}

func NewView(a fyne.App, model *Model) *View {
	v := &View{
		app:   a,
		tasks: make([]string, 0),
		model: model,
	}
	v.w = v.app.NewWindow("ToDo List")
	v.createUI()
	return v
}

func (v *View) createUI() {
	v.addButton = widget.NewButton("Add", func() {
		v.onAddTask(v.entry.Text)
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

func (v *View) Init() error {
	err := v.UpdateTaskList()
	if err != nil {
		return err
	}
	return nil
}

func (v *View) BindAddTask(f func(string)) {
	v.onAddTask = f
}

func (v *View) BindDeleteTask(f func(string)) {
	v.onDeleteTask = f
}

func (v *View) ClearEntry() {
	v.entry.SetText("")
}

func (v *View) UpdateTaskList() error {
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

func (v *View) ShowError(err error) {
	dialog.ShowError(err, v.w)
}

func (v *View) Run() {
	v.w.ShowAndRun()
}
