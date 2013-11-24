package controllers

import (
        "database/sql"
//        "fmt"
        _ "github.com/mattn/go-sqlite3"
        "log"
//        "os"
)

type Film struct {
     Cine 	 string	
     Edo	 string	
     Col	 string	
     CineId	 string	
     CineName	 string	
     Title	 string	
     Rating	 string	
     Language	 string	
     RoomType	 string	
     Date	 string	
     Time	 string	
}

func GetFilms() []Film {
     db, err := sql.Open("sqlite3", "/home/user/Documents/code/github/kartelera.mx/src/prueba/app/controllers/foo.db")
     if err != nil {
     	log.Fatal(err)
     }
     defer db.Close()
     rows, err := db.Query("select cine, edo, col, cineId, cineName, title, rating, language, roomType, date, time from funciones")
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
	 var cineName string
	 var cineId string
	 var col string

     	 rows.Scan(&cine, &edo, &col, &cineId, &cineName, &title, &rating, &language, &roomType, &date, &time)
 	 f := Film{cine, edo, col, cineId, cineName, title, rating, language, roomType, date, time}
	 funciones = append(funciones, f)
     }
     rows.Close()
     return funciones
}
