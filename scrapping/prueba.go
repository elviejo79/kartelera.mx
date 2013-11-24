package main

import (
        "database/sql"
        _ "github.com/mattn/go-sqlite3"
	"fmt"
        "log"
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

func main(){
     GetFilms()
}

func GetFilms() []Film {
     db, err := sql.Open("sqlite3", "./foo.db")
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
	 fmt.Print(f)
	 funciones = append(funciones, f)
     }
     rows.Close()
     return funciones
}

func GetFilmsByEdo(state string) []Film {
     db, err := sql.Open("sqlite3", "./foo.db")
     if err != nil {
     	log.Fatal(err)
     }
     defer db.Close()
     stmt, err := db.Prepare("select cine, edo, col, cineId, cineName, title, rating, language, roomType, date, time from funciones where edo = ?")
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
	 var cineName string
	 var cineId string
	 var col string
	 rows.Scan(&cine, &edo, &col, &cineId, &cineName, &title, &rating, &language, &roomType, &date, &time)
	 f := Film{cine, edo, col, cineId, cineName, title, rating, language, roomType, date, time}
	 funciones = append(funciones, f)
     }
     stmt.Close()
     return funciones
}

