package viewmodel

import (
	"log"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/model"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/view"

	"fyne.io/fyne/v2/data/binding"
)

type ViewModel struct {
	entry binding.String
	model *model.Model
	tasks binding.StringList
	view  *view.View
}

func NewViewModel(m *model.Model, v *view.View) *ViewModel {
	c := &ViewModel{
		entry: binding.NewString(),
		model: m,
		tasks: binding.NewStringList(),
		view:  v,
	}
	return c
}

func (vm *ViewModel) OnAddTask() {
	task, err := vm.entry.Get()
	if err != nil {
		panic(err)
	}
	err = vm.entry.Set("")
	if err != nil {
		panic(err)
	}
	if task == "" {
		return
	}
	err = vm.model.AddTask(task)
	if err != nil {
		log.Fatal(err)
	}
	vm.updateTaskList()
}

func (vm *ViewModel) OnDeleteTask(id int) {
	if id == 0 {
		return
	}
	task, err := vm.tasks.GetValue(id)
	if err != nil {
		log.Fatal(err)
	}
	err = vm.model.DeleteTask(task)
	if err != nil {
		log.Fatal(err)
	}
	vm.updateTaskList()
}

func (vm *ViewModel) updateTaskList() {
	tasks, err := vm.model.ListTasks()
	if err != nil {
		log.Fatal(err)
	}
	err = vm.tasks.Set(tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func (vm *ViewModel) Run() {
	vm.view.InitUI(vm.entry, vm.tasks, vm)
	vm.updateTaskList()
	vm.view.Run()
}
