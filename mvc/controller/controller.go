package controller

import (
	"github.com/ErikKalkoken/fyne-gui-architectures/mvc/model"
	"github.com/ErikKalkoken/fyne-gui-architectures/mvc/view"
)

type Controller struct {
	model *model.Model
	view  *view.View
}

func New(m *model.Model, v *view.View) *Controller {
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
		c.view.ShowError(err)
		return
	}
	c.view.ClearEntry()
	go func() {
		err = c.view.UpdateTaskList()
		if err != nil {
			c.view.ShowError(err)
		}
	}()
}

func (c *Controller) deleteTask(task string) {
	err := c.model.DeleteTask(task)
	if err != nil {
		c.view.ShowError(err)
		return
	}
	go func() {
		err = c.view.UpdateTaskList()
		if err != nil {
			c.view.ShowError(err)
		}
	}()
}

func (c *Controller) Run() {
	c.view.Run()
}
