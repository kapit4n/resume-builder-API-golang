package controllers

import (
	"github.com/revel/revel"
	"rb-rgo/app/models"
	"rb-rgo/app/repository"
)

type Queue []*models.Resume

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) ResumeList() revel.Result {
	return c.RenderJSON(repository.GetResumeRepository().GetResumes())
}

func (c App) ResumeById(id string) revel.Result {
	res := &models.Resume{Id: "1", Title: "Full Stack Developer", Summary: "Summary"}
	return c.RenderJSON(res)
}
