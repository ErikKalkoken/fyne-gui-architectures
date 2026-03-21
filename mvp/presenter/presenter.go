package presenter

import (
	"github.com/ErikKalkoken/fyne-gui-architectures/mvp/model"
)

type View interface {
	ClearEntry()
	ShowError(err error)
	UpdateTaskList([]string)
}

type Presenter struct {
	model *model.Model
	view  View
}

func New(m *model.Model, v View) *Presenter {
	p := &Presenter{
		model: m,
		view:  v,
	}
	return p
}

func (p *Presenter) OnAddTask(task string) {
	err := p.model.AddTask(task)
	if err != nil {
		p.view.ShowError(err)
		return
	}
	p.view.ClearEntry()
	p.updateTaskList()
}

func (p *Presenter) OnDeleteTask(task string) {
	err := p.model.DeleteTask(task)
	if err != nil {
		p.view.ShowError(err)
		return
	}
	p.updateTaskList()

}

func (p *Presenter) updateTaskList() {
	tasks, err := p.model.ListTasks()
	if err != nil {
		p.view.ShowError(err)
		return
	}
	p.view.UpdateTaskList(tasks)
}

func (p *Presenter) Init() {
	p.updateTaskList()
}
