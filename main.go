package main

import (
	"fmt"
	groupietrackers "groupie-tracker/functions"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var currentband groupietrackers.CurrentBand
var pageData groupietrackers.PageData
var artistLoad []groupietrackers.Artist
var artistFiltered []groupietrackers.Artist
var currentID int
var data groupietrackers.ApiData

var FakeCurrentYear int
var FakeCurrentMonth time.Month
var FakeCurrentDay int

func main() {

	FakeCurrentYear, FakeCurrentMonth, FakeCurrentDay = time.Now().Date()
	FakeCurrentYear -= 3
	fmt.Println("Date Simulated :", FakeCurrentDay, FakeCurrentMonth, FakeCurrentYear)

	data = groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))

	go RealtimeData() /*       Permet de récuperer les artists en arrière plan sans ralentir l'affichage de la page         */
	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":8080", nil)
}

func RealtimeData() {
	for { /*       Regenere les données des artistes toutes les minutes        */
		GetArtistXtoY(1, 52, data.Artist)
		time.Sleep(60 * time.Second)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./Login.html"))
	tmpl.Execute(w, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/index.html", "./server/component/sidebar.html"))

	switch r.Method {
	case "GET":
		if r.URL.Query().Get("id") != "" {
			currentID, _ = strconv.Atoi(r.URL.Query().Get("id"))
		}
		pageData.Artists = artistLoad
	case "POST":
		shearchFilter := r.FormValue("shearch")
		artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistLoad)
		pageData.Artists = artistFiltered
	}
	rand.Seed(time.Now().UnixNano())
	pageData.MPageRArtist = []groupietrackers.Artist{}
	mi := 3
	if len(artistLoad) < 3 {
		mi = len(artistLoad)
	}
	for i := 0; i < mi; i++ {
		pageData.MPageRArtist = append(pageData.MPageRArtist, artistLoad[rand.Intn(len(artistLoad))])
	}

	tmpl.Execute(w, pageData)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./server/artist.html", "./server/component/sidebar.html"))

	switch r.Method {
	case "GET":
		if r.URL.Query().Get("id") != "" {
			currentID, _ = strconv.Atoi(r.URL.Query().Get("id"))
		}
		pageData.Artists = artistLoad
	case "POST":
		shearchFilter := r.FormValue("shearch")
		artistFiltered = groupietrackers.GetArtistWithStr(shearchFilter, artistLoad)
		pageData.Artists = artistFiltered
	}

	if currentID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID))
	pageData.Currentband = currentband
	tmpl.Execute(w, pageData)
}

func GetArtistXtoY(x, y int, apiArtist string) {
	for i := x; i <= y; i++ {
		if i > len(artistLoad) {
			artist := groupietrackers.SetArtistInfoData(groupietrackers.GetAPIData(apiArtist + "/" + strconv.Itoa(i)))
			artistLoad = append(artistLoad, artist)
		}
	}
}
