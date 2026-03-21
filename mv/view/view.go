package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/ErikKalkoken/fyne-gui-architectures/mv/model"
)

type View struct {
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

func New(a fyne.App, model *model.Model) *View {
	v := &View{
		app:   a,
		tasks: make([]string, 0),
		model: model,
	}
	v.w = v.app.NewWindow("ToDo List")
	v.makeUI()
	return v
}

func (v *View) makeUI() {
	v.addButton = widget.NewButton("Add", func() {
		go v.addTask(v.entry.Text)
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
			go v.deleteTask(x)
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
			s := v.tasks[id]
			co.(*widget.Label).SetText(s)
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

func (v *View) addTask(task string) {
	err := v.model.AddTask(task)
	if err != nil {
		fyne.Do(func() {
			v.ShowError(err)
		})
	}
	fyne.Do(func() {
		v.ClearEntry()
	})
	go func() {
		err = v.UpdateTaskList()
		if err != nil {
			fyne.Do(func() {
				v.ShowError(err)
			})
		}
	}()
}

func (v *View) deleteTask(task string) {
	err := v.model.DeleteTask(task)
	if err != nil {
		fyne.Do(func() {
			v.ShowError(err)
		})
	}
	go func() {
		err = v.UpdateTaskList()
		if err != nil {
			fyne.Do(func() {
				v.ShowError(err)
			})
		}
	}()
}

func (v *View) Init() error {
	err := v.UpdateTaskList()
	if err != nil {
		return err
	}
	return nil
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
