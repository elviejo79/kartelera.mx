package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index(edo string) revel.Result {
     films := GetTitlesByEdo(edo)
     horarios := GetFilmsByEdo(edo)
     horariosByFilm := make(map [string][]Film)
     for _, horario := range horarios{
     	 horariosByFilm[horario.Title] = append(horariosByFilm[horario.Title], horario)
     }
     return c.Render(films, horariosByFilm)
}

func (c App) Lista(edo string) revel.Result{
     films := GetFilmsByEdo(edo)
     return c.Render(films)
}
