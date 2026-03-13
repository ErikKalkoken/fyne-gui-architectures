package tasks

import (
	"log"
)

type Controller struct {
	model *Model
	view  *View
}

func NewController(m *Model, v *View) *Controller {
	c := &Controller{
		model: m,
		view:  v,
	}
	v.BindAddTask(c.addTask)
	v.BindDeleteTask(c.deleteTask)
	return c
}

func (c *Controller) addTask(task string) {
	err := c.model.AddTask(task)
	if err != nil {
		log.Fatal(err)
	}
	c.view.ClearEntry()
	go func() {
		err = c.view.UpdateTaskList()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (c *Controller) deleteTask(task string) {
	err := c.model.DeleteTask(task)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = c.view.UpdateTaskList()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (c *Controller) Run() {
	c.view.Run()
}
