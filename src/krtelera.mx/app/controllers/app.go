package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
     films := GetFilms()
     return c.Render(films)
}

func (c App) Hello(myName string) revel.Result{
     return c.Render(myName)
}
