package tasks

import (
	"log"
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

func (p *Presenter) HandleAddTask() {
	task := p.view.EntryText()
	err := p.model.AddTask(task)
	if err != nil {
		log.Fatal(err)
	}
	p.view.ClearEntry()
	p.updateTaskList()
}

func (p *Presenter) HandleDeleteTask() {
	task := p.view.SelectedText()
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
	p.view.UpdateTaskList(tasks)
}

func (p *Presenter) Run() {
	p.view.InitUI(p)
	p.updateTaskList()
	p.view.Run()
}
