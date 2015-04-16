package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

type Person struct {
	Name string
	Age  int
}

func (c App) Index() revel.Result {
	title := "Index"
	name := "Page 1"
	return c.Render(title, name)
}

func (c App) Name() revel.Result {
	var person Person
	c.Params.Bind(&person, "person")
	return c.Render(person)
}
