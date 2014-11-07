package controllers

import (
	"github.com/JonPulfer/rbsc/app/models"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.RenderArgs["articles"] = models.AllArticles()
	return c.Render()
}
