package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Game struct {
	ID        string     `json:"id"`
	Isbn      string     `json:"isbn"`
	Title     string     `json:"title"`
	Developer *Developer `json:"developer"`
}

type Developer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var games []Game

func main() {
	r := mux.NewRouter()

	games = append(games, Game{ID: "1", Isbn: "1710991059", Title: "GTA", Developer: &Developer{FirstName: "Rockstar", LastName: "Games"}})
	games = append(games, Game{ID: "2", Isbn: "1710991049", Title: "Valorant", Developer: &Developer{FirstName: "Riot", LastName: "Games"}})

	r.HandleFunc("/games", getGames).Methods("GET")
	r.HandleFunc("/games/{id}", getGame).Methods(("GET"))
	r.HandleFunc("/games", createGame).Methods("POST")
	r.HandleFunc("/games/{id}", updateGame).Methods("PUT")
	r.HandleFunc("/games/{id}", deleteGame).Methods("DELETE")

	fmt.Print("Serving starting at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

// Get all The Games
func getGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}

// Delete The specefic Movie
func deleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, game := range games {
		if game.ID == params["id"] {
			games = append(games[:index], games[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(games)
}

//Get one Game

func getGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for _, game := range games {
		if game.ID == param["id"] {
			json.NewEncoder(w).Encode(game)
			return
		}
	}
}

// create game
func createGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var game Game
	_ = json.NewDecoder(r.Body).Decode(&game)
	game.ID = strconv.Itoa(rand.Intn(10000000000))
	games = append(games, game)
	json.NewEncoder(w).Encode(game)
}

// update Game
func updateGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, game := range games {
		if game.ID == param["id"] {
			games = append(games[:index], games[index+1:]...)
			var game Game
			_ = json.NewDecoder(r.Body).Decode(&game)
			game.ID = param["id"]
			games = append(games, game)
			json.NewEncoder(w).Encode(game)
			return
		}
	}
}
