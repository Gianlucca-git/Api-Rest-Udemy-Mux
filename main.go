package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/",funtion )
	router.HandleFunc( "/peliculas", MovieList)
	router.HandleFunc( "/peliculas/{id}", MovieShow)

	// podemos poner las rutas que me den la gana

	server := http.ListenAndServe(":3128",router)
	log.Fatal(server)
}
func funtion (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "HOLA MUNDO DESDE EL SERVIDEO WEB GO 2 ")
}

func MovieList (w http.ResponseWriter, r *http.Request){

	type Movie struct {
		Name     string
		Gener    string
		Year     int
		Director string
	}

	type Movies []Movie

	movies :=  Movies{
		Movie{"kunfu panda", "animada", 2015, "Gian"},
		Movie{"shrek 3", "animada", 2020, "Lucca"},
	}

	//fmt.Fprintf(w, "HOLA MUNDO DESDE MOVIE LIST ")
	json.NewEncoder(w).Encode(movies)
}

func MovieShow (w http.ResponseWriter, r *http.Request){

	//recoger parametros URL
	params := mux.Vars(r)
	movie_id := params["id"]


	fmt.Fprintf(w, "HOLA MUNDO DESDE MOVIE LIST %s ", movie_id)
}


