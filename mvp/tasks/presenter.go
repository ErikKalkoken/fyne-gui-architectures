package tasks

import (
	"log"

	"fyne.io/fyne/v2"
)

type Presenter struct {
	model *Model
	view  *View
}

func NewPresenter(m *Model, v *View) *Presenter {
	c := &Presenter{
		model: m,
		view:  v,
	}
	return c
}

func (p *Presenter) OnAddTask(task string) {
	err := p.model.AddTask(task)
	if err != nil {
		log.Fatal(err)
	}
	fyne.Do(func() {
		p.view.ClearEntry()
	})
	p.updateTaskList()
}

func (p *Presenter) OnDeleteTask(task string) {
	err := p.model.DeleteTask(task)
	if err != nil {
		log.Fatal(err)
	}
	p.updateTaskList()

}

func (p *Presenter) updateTaskList() {
	tasks, err := p.model.ListTasks()
	if err != nil {
		log.Fatal(err)
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
