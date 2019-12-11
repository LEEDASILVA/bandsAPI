package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate string   `json:"date"`
}

type Bands struct {
	Index []Artist `json:"index"`
}

type Locations struct {
	ID    int      `json:"id"`
	Locat []string `json:"locations"`
}

type Locs struct {
	Index []Locations `json:"index"`
}

type Dates struct {
	ID  int      `json:"id"`
	Dat []string `json:"dates"`
}

type Dats struct {
	Index []Dates `json:"index"`
}

type DatsLocs struct {
	ID       int                 `json:"id"`
	DatsLocs map[string][]string `json:"datesLocations"`
}

type DatLoc struct {
	Index []DatsLocs `json:"index"`
}

var bands Bands
var artists []Artist

var locs Locs
var locations []Locations

var dats Dats
var dates []Dates

var datslocs DatsLocs
var datloc DatLoc

//Get All Artists
func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bands)
}

//get All dates and locations
func getRelation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datloc)
}

//Get One Artists
func getArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through artists and find with name
	for _, item := range artists {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Artist{})
}

func getJSON(jsonfile, types string) {
	jsonFile, err := os.Open(jsonfile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %v\n", jsonfile)

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	switch types {
	case "artists":
		var b Bands

		errors := json.Unmarshal(byteValue, &b)
		if errors != nil {
			fmt.Println(errors)
		}

		for i := 0; i < len(b.Index); i++ {
			elem := b.Index[i]
			artists = append(artists, elem)
		}
		bands = Bands{Index: artists}

	case "locations":
		var l Locs
		errors := json.Unmarshal(byteValue, &l)

		if errors != nil {
			fmt.Println(errors)
		}
		for i := 0; i < len(l.Index); i++ {
			elem := l.Index[i]
			locations = append(locations, elem)
		}
		locs = Locs{Index: locations}

	case "dates":
		var d Dats
		errors := json.Unmarshal(byteValue, &d)

		if errors != nil {
			fmt.Println(errors)
		}
		for i := 0; i < len(d.Index); i++ {
			elem := d.Index[i]
			dates = append(dates, elem)
		}
		dats = Dats{Index: dates}
	}

	jsonFile.Close()
}

func joinStructs(d Dats, l Locs) {
	for i, dat := range d.Index {
		arrDates := getDates(dat)
		for j, x := range l.Index[i].Locat {

			//x -> cada localização de um certo id i
			//arrDates -> todas as datas dum um certo id i
			datslocs.DatsLocs[x] = arrDates[j]
		}
		datslocs.ID = i + 1
		datloc.Index = append(datloc.Index, datslocs)
		datslocs.DatsLocs = map[string][]string{}
	}

}

func getDates(d Dates) [][]string {

	var helper [][]string
	var helper2 []string
	aster := false
	for i, dat := range d.Dat {
		if dat[0] == '*' && aster == false {
			helper2 = append(helper2, dat[1:])
			aster = true
		} else if dat[0] == '*' && aster == true {
			helper = append(helper, helper2)
			helper2 = []string{}
			helper2 = append(helper2, dat[1:])
		} else if dat[0] != '*' && aster == true {
			helper2 = append(helper2, dat)
		}

		if i == len(d.Dat)-1 {
			helper = append(helper, helper2)
		}

	}
	return helper
}

// func createArtist(w http.ResponseWriter, r *http.Request) {

// }

// func updateArtist(w http.ResponseWriter, r *http.Request) {

// }

// func deleteArtist(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	r := mux.NewRouter()

	getJSON("../data/artists.json", "artists")
	getJSON("../data/locations.json", "locations")
	getJSON("../data/dates.json", "dates")

	if datslocs.DatsLocs == nil {
		datslocs.DatsLocs = make(map[string][]string)
	}
	joinStructs(dats, locs)

	//Route Handlers / Endpoints
	r.HandleFunc("/api/artists", getArtists).Methods("GET")
	r.HandleFunc("/api/relation", getRelation).Methods("GET")

	r.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")

	// r.HandleFunc("/api/artists", createArtist).Methods("POST")
	// r.HandleFunc("/api/artists/{id}", updateArtist).Methods("PUT")
	// r.HandleFunc("/api/artists/{id}", deleteArtist).Methods("DELETE")

	fmt.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
