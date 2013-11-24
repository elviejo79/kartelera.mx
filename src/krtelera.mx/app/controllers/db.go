package controllers

import (
        "database/sql"
//        "fmt"
        _ "github.com/mattn/go-sqlite3"
        "log"
//        "os"
)

const dataBase string = "/home/user/Documents/code/github/kartelera.mx/scrapping/foo.db"

type Film struct {
     Cine 	 string	
     Edo	 string	
     Col	 string	
     CineId	 string	
     CineName	 string	
     Title	 string	
     Img	 string
     Rating	 string	
     Language	 string	
     RoomType	 string	
     Date	 string	
     Time	 string	
}

type Titles struct{
     Title string
     Img   string
}


func GetFilms() []Film {
     db, err := sql.Open("sqlite3", dataBase)
     if err != nil {
     	log.Fatal(err)
     }
     defer db.Close()
     rows, err := db.Query("select cine, edo, col, cineId, cineName, title, img, rating, language, roomType, date, time from funciones")
     if err != nil {
          log.Fatal(err)
     }

     defer rows.Close()
     var funciones []Film

     for rows.Next(){
     	 var cine string
	 var edo string
	 var time string
	 var date string
	 var roomType string
	 var language string
	 var rating string
	 var title string
	 var img string
	 var cineName string
	 var cineId string
	 var col string

     	 rows.Scan(&cine, &edo, &col, &cineId, &cineName, &title, &img, &rating, &language, &roomType, &date, &time)
 	 f := Film{cine, edo, col, cineId, cineName, title, img, rating, language, roomType, date, time}
	 funciones = append(funciones, f)
     }
     rows.Close()
     return funciones
}

func GetFilmsByEdo(state string) []Film {
     db, err := sql.Open("sqlite3", dataBase)
     if err != nil {
     	log.Fatal(err)
     }
     defer db.Close()
     stmt, err := db.Prepare("select cine, edo, col, cineId, cineName, title, img, rating, language, roomType, date, time from funciones where edo = ? order by title, time")
     if err != nil {
          log.Fatal(err)
     }
     defer stmt.Close()
     var funciones []Film
     rows, err := stmt.Query(state)
          if err != nil {
          log.Fatal(err)
     }
     for rows.Next(){
	 var cine string
	 var edo string
	 var time string
	 var date string
	 var roomType string
	 var language string
	 var rating string
	 var title string
	 var img string
	 var cineName string
	 var cineId string
	 var col string
	 rows.Scan(&cine, &edo, &col, &cineId, &cineName, &title, &img, &rating, &language, &roomType, &date, &time)
	 f := Film{cine, edo, col, cineId, cineName, title, img, rating, language, roomType, date, time}
	 funciones = append(funciones, f)
     }
     stmt.Close()
     return funciones
}


func GetTitlesByEdo(state string) []Titles {
     db, err := sql.Open("sqlite3", dataBase)
     if err != nil {
     	log.Fatal(err)
     }
     defer db.Close()
     stmt, err := db.Prepare("select title, img from funciones where edo = ? group by title")
     if err != nil {
          log.Fatal(err)
     }
     defer stmt.Close()
     var funciones []Titles
     rows, err := stmt.Query(state)
          if err != nil {
          log.Fatal(err)
     }
     for rows.Next(){
	 var title string
	 var img string
	 rows.Scan(&title, &img)
	 f := Titles{title, img}
	 funciones = append(funciones, f)
     }
     stmt.Close()
     return funciones
}

