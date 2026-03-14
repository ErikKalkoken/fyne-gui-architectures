package tasks

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type View struct {
	addButton    *widget.Button
	app          fyne.App
	deleteButton *widget.Button
	entry        *widget.Entry
	taskList     *widget.List
	w            fyne.Window
}

func NewView(a fyne.App, vm *ViewModel) *View {
	v := &View{
		app: a,
		w:   a.NewWindow("ToDo List"),
	}

	v.addButton = widget.NewButton("Add", func() {
		go vm.OnAddTask()
	})
	v.addButton.Disable()

	v.entry = widget.NewEntryWithData(vm.Entry)
	v.entry.PlaceHolder = "Name of new task"
	v.entry.OnChanged = func(s string) {
		if len(s) > 0 {
			v.addButton.Enable()
		} else {
			v.addButton.Disable()
		}
	}

	var selectedID int
	v.deleteButton = widget.NewButton("Delete", func() {
		go vm.OnDeleteTask(selectedID)
	})
	v.deleteButton.Disable()

	v.taskList = widget.NewListWithData(
		vm.Tasks,
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			v, err := di.(binding.String).Get() // FIXME: Workaround for the out of bounds issue
			if err != nil {
				panic(err)
			}
			co.(*widget.Label).SetText(v)
		},
	)
	v.taskList.OnSelected = func(id widget.ListItemID) {
		selectedID = id
		v.deleteButton.Enable()
	}
	v.taskList.OnUnselected = func(id widget.ListItemID) {
		selectedID = 0
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
	return v
}

func (v *View) Run() {
	v.w.ShowAndRun()
}
