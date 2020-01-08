package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Bands struct {
	Index []Artist `json:"index"`
}
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locs struct {
	Index []Locations `json:"index"`
}
type Locations struct {
	ID    int      `json:"id"`
	Locat []string `json:"locations"`
}

type Dats struct {
	Index []Dates `json:"index"`
}
type Dates struct {
	ID  int      `json:"id"`
	Dat []string `json:"dates"`
}

type DatLoc struct {
	Index []DatsLocs `json:"index"`
}
type DatsLocs struct {
	ID       int                 `json:"id"`
	DatsLocs map[string][]string `json:"datesLocations"`
}

var bands Bands
var artists []Artist

var locs Locs
var locations []Locations

var dats Dats
var dates []Dates

var datslocs DatsLocs
var datloc DatLoc

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// all Get handlers
func getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(artists)
	json.NewEncoder(w).Encode(artists)
}

func getLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locs)
}

func getDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dats)
}

func getRelations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datloc)
}

func getLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type a struct {
		A string `json:"artists"`
		L string `json:"locations"`
		D string `json:"dates"`
		R string `json:"relation"`
	}
	res := a{}
	str := `{ "artists": "https://groupietrackers.herokuapp.com/api/artists",
	"locations": "https://groupietrackers.herokuapp.com/api/locations",
	"dates": "https://groupietrackers.herokuapp.com/api/dates",
	"relation": "https://groupietrackers.herokuapp.com/api/relation" }`
	json.Unmarshal([]byte(str), &res)
	json.NewEncoder(w).Encode(res)
}

// all Get handlers using url querys
func getArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// loop through artists and find with name
	for _, item := range artists {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Artist{})
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range locations {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Artist{})
}

func getDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range dates {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Artist{})
}

func getRelation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range datloc.Index {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Artist{})
}

func getImages(w http.ResponseWriter, r *http.Request) {
	buffer := new(bytes.Buffer)
	params := mux.Vars(r)

	image := "../images/" + params["image"]
	if err := jpeg.Encode(buffer, loadImage(image), nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

//decode the image and returns it
func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	handleError(err)

	defer f.Close()

	img, err := jpeg.Decode(f)
	handleError(err)
	return img
}

func getJSON(jsonfile string) {
	jsonFile, err := os.Open(jsonfile)
	handleError(err)

	// read our opened xml/json as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	handleError(err)
	types := strings.TrimPrefix(
		strings.TrimSuffix(jsonfile, ".json"), "../data/")

	switch types {
	case "artists":
		var b Bands
		fmt.Println(byteValue)
		err := json.Unmarshal(byteValue, &b)
		handleError(err)

		for i := 0; i < len(b.Index); i++ {
			elem := b.Index[i]
			artists = append(artists, elem)
		}
		bands = Bands{Index: artists}

	case "locations":
		var l Locs
		err := json.Unmarshal(byteValue, &l)
		handleError(err)

		for i := 0; i < len(l.Index); i++ {
			elem := l.Index[i]
			locations = append(locations, elem)
		}
		locs = Locs{Index: locations}

	case "dates":
		var d Dats
		err := json.Unmarshal(byteValue, &d)
		handleError(err)

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
		arrDates := locationNdates(dat)
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

func locationNdates(d Dates) [][]string {
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

func main() {
	r := mux.NewRouter()
	//r.Handle("/", http.FileServer(http.Dir(".")))

	getJSON("data/artists.json")
	getJSON("data/locations.json")
	getJSON("data/dates.json")

	if datslocs.DatsLocs == nil {
		datslocs.DatsLocs = make(map[string][]string)
	}
	joinStructs(dats, locs)

	//Route Handlers / Endpoints
	r.HandleFunc("/api", getLink).Methods("GET")
	r.HandleFunc("/api/artists", getArtists).Methods("GET")
	r.HandleFunc("/api/locations", getLocations).Methods("GET")
	r.HandleFunc("/api/dates", getDates).Methods("GET")
	r.HandleFunc("/api/relation", getRelations).Methods("GET")

	r.HandleFunc("/api/relation/{id}", getRelation).Methods("GET")
	r.HandleFunc("/api/dates/{id}", getDate).Methods("GET")
	r.HandleFunc("/api/locations/{id}", getLocation).Methods("GET")
	r.HandleFunc("/api/artists/{id}", getArtist).Methods("GET")
	r.HandleFunc("/api/images/{image}", getImages).Methods("GET")

	port := GetPort()
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("Server running on port%v", port)
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
