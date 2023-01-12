package main

import (
	groupietrackers "groupie-tracker/functions"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var currentband groupietrackers.CurrentBand
var pageData groupietrackers.PageData
var artistLoad []groupietrackers.Artist
var artistFiltered []groupietrackers.Artist
var currentID int
var data groupietrackers.ApiData

func main() {

	data = groupietrackers.SetGlobalData(groupietrackers.GetAPIData("https://groupietrackers.herokuapp.com/api"))

	go RealtimeData() /*       Permet de récuperer les artists en arrière plan sans ralentir l'affichage de la page         */
	fs := http.FileServer(http.Dir("./server"))
	http.Handle("/server/", http.StripPrefix("/server/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/artist", ArtistHandler)
	http.ListenAndServe(":8080", nil)
}

func RealtimeData() {
	for { /*       Regenere les données des artistes toutes les minutes        */
		GetArtistXtoY(1, 52, data.Artist)
		time.Sleep(60 * time.Second)
	}
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
		GetArtistWithStr(shearchFilter)
		pageData.Artists = artistFiltered
	}
	rand.Seed(time.Now().UnixNano())
	pageData.MPageRArtist = []int{}
	pageData.MPageRArtist = append(pageData.MPageRArtist, rand.Intn(len(pageData.Artists)))

	//currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID))
	//pageData.Currentband = currentband

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
		GetArtistWithStr(shearchFilter)
		pageData.Artists = artistFiltered
	}

	if currentID == 0 {
		http.Redirect(w, r, "/", http.StatusFound)
	}

	currentband = groupietrackers.UpdateCurrentBand(data.Artist + "/" + strconv.Itoa(currentID))
	pageData.Currentband = currentband
	tmpl.Execute(w, pageData)
}

func GetArtistWithStr(shearchFilter string) {
	artistFiltered = []groupietrackers.Artist{}
	for _, artist := range artistLoad {
		if strings.Contains(TurnStringToShearch(artist.Name), TurnStringToShearch(shearchFilter)) {
			artistFiltered = append(artistFiltered, artist)
		}
	}
}

func TurnStringToShearch(str string) string {
	/*       Turn :  fdsfKJHJUGKHLJ dsf ezrtf _è-'4941 into : fdsfkjhjugkhljdsfezrtf_è-'4941*/
	var nstr string
	for _, car := range str {
		switch {
		case 65 <= car && car <= 90:
			nstr = nstr + string(car+32)
		case car == 32:
			continue
		default:
			nstr = nstr + string(car)
		}
	}
	return nstr
}

func GetArtistXtoY(x, y int, apiArtist string) {
	for i := x; i <= y; i++ {
		if i > len(artistLoad) {
			artist := groupietrackers.SetArtistInfoData(groupietrackers.GetAPIData(apiArtist + "/" + strconv.Itoa(i)))
			artistLoad = append(artistLoad, artist)
		}
	}
}
