package controllers

import (
	"github.com/revel/revel"
	"rb-rgo/app/models"
)

type Queue []*models.Resume

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) ResumeList() revel.Result {
	var q Queue
	resume := &models.Resume{Id: "1", Title: "Full Stack GO Developer", Summary: "Summary"}
	q = append(q, resume)
	return c.RenderJSON(q)
}

func (c App) ResumeById() revel.Result {
	res := &models.Resume{Id: "1", Title: "Full Stack Developer", Summary: "Summary"}
	return c.RenderJSON(res)
}
