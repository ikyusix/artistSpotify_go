package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handlerArtist "artistSpotify/artist/handler"
	repoArtist "artistSpotify/artist/repo"
	usecaseArtist "artistSpotify/artist/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	port := "8080"
	conStr := "root:root@tcp(127.0.0.1)/spotify_db"

	db, err := sql.Open("mysql", conStr)
	defer db.Close()

	r := mux.NewRouter()

	artistRepo := repoArtist.CreateArtistRepoMysqlImpl(db)
	artistUsecase := usecaseArtist.CreateArtistUsecaseImpl(artistRepo)

	handlerArtist.CreateArtistHandler(r, artistUsecase)

	fmt.Println("Starting web server at http://localhost:8080")
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
