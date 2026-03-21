package viewmodel

import (
	"fyne.io/fyne/v2/data/binding"

	"github.com/ErikKalkoken/fyne-gui-architectures/mvvm/model"
)

type ViewModel struct {
	Entry binding.String
	Tasks binding.StringList
	Error binding.Item[error]

	model *model.Model
}

func New(m *model.Model) *ViewModel {
	vm := &ViewModel{
		Entry: binding.NewString(),
		Error: binding.NewItem(func(a, b error) bool {
			return a == b
		}),
		Tasks: binding.NewStringList(),

		model: m,
	}
	return vm
}

func (vm *ViewModel) Init() error {
	return vm.updateTaskList()
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
	err = vm.updateTaskList()
	if err != nil {
		vm.Error.Set(err)
		return
	}
}

func (vm *ViewModel) OnDeleteTask(id int) {
	if id == 0 {
		return
	}
	task, err := vm.Tasks.GetValue(id)
	if err != nil {
		vm.Error.Set(err)
		return
	}
	err = vm.model.DeleteTask(task)
	if err != nil {
		vm.Error.Set(err)
		return
	}
	err = vm.updateTaskList()
	if err != nil {
		vm.Error.Set(err)
		return
	}
}

func (vm *ViewModel) updateTaskList() error {
	tasks, err := vm.model.ListTasks()
	if err != nil {
		return err
	}
	err = vm.Tasks.Set(tasks)
	if err != nil {
		return err
	}
	return nil
}
