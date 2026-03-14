package main

import (
	"fyne.io/fyne/v2"
)

type Presenter struct {
	model *Model
	view  *View
}

func NewPresenter(m *Model, v *View) *Presenter {
	p := &Presenter{
		model: m,
		view:  v,
	}
	return p
}

func (p *Presenter) OnAddTask(task string) {
	err := p.model.AddTask(task)
	if err != nil {
		fyne.Do(func() {
			p.view.ShowError(err)
		})
	}
	fyne.Do(func() {
		p.view.ClearEntry()
	})
	p.updateTaskList()
}

func (p *Presenter) OnDeleteTask(task string) {
	err := p.model.DeleteTask(task)
	if err != nil {
		fyne.Do(func() {
			p.view.ShowError(err)
		})
	}
	p.updateTaskList()

}

func (p *Presenter) updateTaskList() {
	tasks, err := p.model.ListTasks()
	if err != nil {
		fyne.Do(func() {
			p.view.ShowError(err)
		})
	}
	fyne.Do(func() {
		p.view.UpdateTaskList(tasks)
	})
}

func (p *Presenter) Run() {
	p.view.InitUI(p)
	p.updateTaskList()
	p.view.Run()
}
