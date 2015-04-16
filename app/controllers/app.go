package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	title := "Index"
	name := "Page 1"
	return c.Render(title, name)
}
