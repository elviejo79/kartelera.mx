package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"./cinemex"
	"./cinepolis"
	"bitbucket.org/kardianos/osext"
)

func main() {
	filename, _ := osext.ExecutableFolder()
	cines_db := filename + "funciones.db"
	
	fmt.Println(cines_db)

	os.Remove(cines_db)

	db, err := sql.Open("sqlite3",cines_db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := `
	create table funciones (id integer not null primary key autoincrement, cine text, edo text, col text, cineId text, cineName text, title text, img text, rating text, language text, roomType text, date text, time text);
	`
	_, err = db.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into funciones(cine , edo , col , cineId, cineName , title , img, rating , language , roomType , date , time) values(? , ? , ? , ?, ? , ? , ? , ? , ? , ? , ?, ?)")
		if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _,screening := range cinemex.Screenings() {
		fmt.Print("mx.")
		_,err := stmt.Exec(screening["cine"], screening["edo"] , screening["col"] , screening["cineId"], screening["cineName"] , screening["title"], screening["img"] , screening["rating"] , screening["language"] , screening["roomType"] , screening["date"] , screening["time"])
		if err != nil {
			fmt.Println(err)
		}

	}

	for _,screening := range cinepolis.Screenings() {
		fmt.Print("pl.")
		_,err := stmt.Exec(screening["cine"], screening["edo"] , screening["col"] , screening["cineId"], screening["cineName"] , screening["title"] , screening["img"], screening["rating"] , screening["language"] , screening["roomType"] , screening["date"] , screening["time"])
		if err != nil {
			fmt.Println(err)
		}

	}

	tx.Commit()

/*	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
	rows.Close()

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
*/
}
