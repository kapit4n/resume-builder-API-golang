package controllers

import (
	"github.com/revel/revel"
	"rb-rgo/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Resume() revel.Result {
	res := &models.Resume{Id: "1", Title: "Full Stack Developer", Summary: "Summary"}
	return c.RenderJSON(res)
}
