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
	introText := "RBSC offers its members sailing on the Norfolk Broads 52 weeks of the year. The Club welcomes new members, whether expert sailors or complete novices, young or old! We have a wide range of dinghy classes including:- Moth, Enterprise, Laser, Mirror, Optimist, Phantom, Topper, Solo, Wayfarer, Wanderer. The club also has a selection of boats available for hire to members."
	c.RenderArgs["introText"] = introText
	return c.Render()
}

func (c App) Articles() revel.Result {
	return c.RenderJson(models.AllArticles())
}

func (c App) Article(id int) revel.Result {
	return c.RenderJson(models.ArticleById(id))
}
