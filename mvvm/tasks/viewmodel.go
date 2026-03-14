package tasks

import (
	"log"

	"fyne.io/fyne/v2/data/binding"
)

type ViewModel struct {
	Entry binding.String
	Tasks binding.StringList

	model *Model
}

func NewViewModel(m *Model) *ViewModel {
	vm := &ViewModel{
		Entry: binding.NewString(),
		model: m,
		Tasks: binding.NewStringList(),
	}
	vm.updateTaskList()
	return vm
}

func (vm *ViewModel) OnAddTask() {
	task, err := vm.Entry.Get()
	if err != nil {
		panic(err)
	}
	err = vm.Entry.Set("")
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
	task, err := vm.Tasks.GetValue(id)
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
	err = vm.Tasks.Set(tasks)
	if err != nil {
		log.Fatal(err)
	}
}
