package tasks

import (
	"log"

	"fyne.io/fyne/v2/data/binding"
)

type ViewModel struct {
	Entry binding.String
	Tasks binding.StringList
	Error binding.Item[error]

	model *Model
}

func NewViewModel(m *Model) *ViewModel {
	vm := &ViewModel{
		Entry: binding.NewString(),
		model: m,
		Tasks: binding.NewStringList(),
		Error: binding.NewItem(func(a, b error) bool {
			return a == b
		}),
	}
	vm.updateTaskList()
	return vm
}

func (vm *ViewModel) OnAddTask() {
	task, err := vm.Entry.Get()
	if err != nil {
		vm.Error.Set(err)
		return
	}
	err = vm.Entry.Set("")
	if err != nil {
		vm.Error.Set(err)
		return
	}
	if task == "" {
		return
	}
	err = vm.model.AddTask(task)
	if err != nil {
		vm.Error.Set(err)
		return
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
